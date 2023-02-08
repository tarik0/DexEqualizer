const argv = require("node:process").argv;
const ethers = require("ethers");
const Web3ProvidersWs = require("web3-providers-ws")
const fs = require("fs");
const readline = require("readline");
const multicall = require("@morpho-labs/ethers-multicall");
const request = require("request");
const http = require("http");

// Event ids.
const TransferEventId = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef";
const SwapEventId = "0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822";

// Maximum token count.
const MaxTokenCount = 3000;
const ResetInterval = 86400 * 1000;

// zeroPad allows the user to pad number with leading zeros.
const zeroPad = (num, places) => String(num).padStart(places, '0')

// loadRouters loads the router addresses from file.
async function loadRouters(chainId) {
    const fileStream = fs.createReadStream(`./chains/${chainId}/routers.txt`);

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });
    // Note: we use the crlfDelay option to recognize all instances of CR LF
    // ('\r\n') in input.txt as a single line break.

    // Iterate over lines.
    let addresses = [];
    for await (const line of rl) {
        // Each line in input.txt will be successively available here as `line`.
        let addressStr = line.split(" ")[0].trim()
        if (ethers.utils.isAddress(addressStr)) addresses.push(ethers.utils.getAddress(addressStr))
    }

    return addresses
}

// getTokenFee retrieves token fees from honeypot.is.
function getTokenFee(address, chain="bsc2") {
    let url = `https://aywt3wreda.execute-api.eu-west-1.amazonaws.com/default/IsHoneypot?chain=${chain}&token=${address}`
    return new Promise((resolve, reject) => {
        request(url, { json: true }, (error, res, body) => {
            // Reject if any error.
            if (error || res.statusCode !== 200) {
                reject(error || `status code is ${res.statusCode}`);
                return;
            }

            resolve(body)
        });
    })
}

// main is the main entry point.
async function main() {
    // Check RPC.
    const rpcUrl = argv.length > 1 ? argv[2] : undefined;
    if (rpcUrl === undefined) {
        console.error("No RPC found in the arguments.")
        return
    }

    // New websocket provider.
    console.log("Connecting to provider:", rpcUrl)
    const provider =  new ethers.providers.Web3Provider(
        new Web3ProvidersWs(rpcUrl, {
            clientConfig: {
                keepalive: true,
                keepaliveInterval: 60000, // ms
            },
            // Enable auto reconnection
            reconnect: {
                auto: true,
                delay: 5000, // ms
                maxAttempts: 5,
                onTimeout: false
            }
        }),
    )

    // New multicall provider.
    const multicallProvider = new multicall.EthersMulticall(provider);

    // Get chain id.
    const network = await provider.getNetwork();
    const chainId = network.chainId;

    // Get routers.
    let routerAddresses = await loadRouters(chainId)

    // The latest block details.
    let lastBlockNum = 0;
    let lastResetDate = 0;
    let tokenSymbols = new Map();
    let tokenUsageMap = new Map();
    let sortedUsageMap = new Map();

    // Sort the usages per 1 min.
    setInterval(async () => {
        // Get token fee infos.
        let tokenInfos = await Promise.all([...tokenUsageMap.entries()].map(async (val) => {
            let [address, usage] = val;

            // Get the fee info.
            let feeInfo;
            try {
                feeInfo = await getTokenFee(address);
            } catch (e) {
                console.error(`Unable to get token fee for ${address}: ${e}`)
                tokenUsageMap.delete(address)
                return;
            }

            return { feeInfo: feeInfo, address: address, usage: usage }
        }));

        // Filter the token fees.
        tokenInfos = tokenInfos.filter((info) => {
            return (
                info !== undefined &&
                info.feeInfo.IsHoneypot === false &&
                info.feeInfo.Error === null &&
                info.feeInfo.BuyTax === 0 &&
                info.feeInfo.SellTax === 0 &&
                info.feeInfo.BuyGas < 160000 &&
                info.feeInfo.SellGas < info.feeInfo.BuyGas
            )
        })

        let zip = (a1, a2) => a1.map((x, i) => [x, a2[i]]);

        // Get token symbols.
        try {
            let symbols = await Promise.all(tokenInfos.map(tokenInfo => {
                let contract = new ethers.Contract(tokenInfo.address, ["function symbol() view returns (string)"])
                return multicallProvider.wrap(contract).symbol();
            }))

            let tokenAddrs = symbols.map((_, i) => tokenInfos[i].address);
            tokenSymbols = new Map(zip(tokenAddrs, symbols))
        } catch (e) {
            console.error(`Unable to get token symbols:`, e)
            return
        }

        // Sort tokens.
        sortedUsageMap = new Map([...tokenUsageMap.entries()]
            .filter(tmp => {
                // Skip the ones without any symbol.
                let [address,] = tmp;
                return tokenSymbols.has(address);
            })
            .filter(tmp => {
                let [_, info] = tmp;
                return info.pairs.length > 1
            })
            .sort((a, b) => {
                return b[1].usage - a[1].usage
            }));

        lastResetDate = Date.now();

        // Clear the console.
        console.clear()
        console.log("Top 10 Hot Tokens (1D) (Trading Volume)")
        console.log(`(Last Update #${lastBlockNum}, ${sortedUsageMap.size} pairs.)`)
        console.log()

        // Print sorted.
        let addresses = Array.from(sortedUsageMap.keys());
        for (let i = 0; i < addresses.length && i < 10; i++) {
            let val = sortedUsageMap.get(addresses[i]);
            let symbol = tokenSymbols.get(addresses[i]);
            console.log(`${zeroPad(i, 2)}) ${addresses[i]} - (Pairs: ${val.pairs.length}) - ${symbol}`)
        }
    }, 60 * 1000)

    // Subscribe to new blocks.
    console.log("Connected to the provider.")
    provider.on("block", async (blockNum) => {
        // Get block transactions.
        let blockWithTxes;
        try {
            blockWithTxes = await provider.getBlockWithTransactions(blockNum);
        } catch (e) {
            console.error("Unable to get block transactions:", e);
            return;
        }

        // Get block logs.
        let blockLogs;
        try {
            blockLogs = await provider.getLogs({
                fromBlock: blockNum,
                toBlock: blockNum
            })
        } catch (e) {
            console.error("Unable to get block logs:", e);
            return;
        }

        // Filter the router transactions.
        let routerTxes = blockWithTxes.transactions.filter(tx => routerAddresses.includes(tx.to))
        let routerTxHashes = routerTxes.map(tx => tx.hash)

        // Filter the transfer logs.
        let transferLogs = blockLogs.filter(log => log.topics.includes(TransferEventId) && routerTxHashes.includes(log.transactionHash));
        let pairContracts = blockLogs.filter(log => log.topics.includes(SwapEventId)).map(log => log.address);

        // Increase the counter for all tokens.
        await Promise.all(transferLogs.map(async (transferLog, i) => {
            // Check if token is a pair.
            if (pairContracts.includes(transferLog.address)) return;

            // Get the previous usage info.
            let prevUsageInfo = tokenUsageMap.get(transferLog.address) || { pairs: [], usage: 0 };

            // Check if we already have too many tokens.
            if (prevUsageInfo.usage === 0 && tokenUsageMap.size > MaxTokenCount) return;

            // Find the swap logs in the transaction.
            let swapLogsInTx = blockLogs
                .filter(log => log.topics.includes(SwapEventId) && log.transactionHash === transferLog.transactionHash)
                .filter(log => log.logIndex > transferLog.logIndex);

            // Find the pairs.
            let pairsInTx = swapLogsInTx.map(log => log.address);

            // Remove duplicates.
            pairsInTx = pairsInTx.filter(function(elem, pos) {
                return pairsInTx.indexOf(elem) === pos;
            })

            // Generate pair token calls.
            for (let i = 0; i < pairsInTx.length; i++) {
                let contract = new ethers.Contract(pairsInTx[i], [
                    "function token0() external view returns (address)",
                    "function token1() external view returns (address)"
                ]);

                try {
                    let [token0, token1] = await Promise.all([
                        multicallProvider.wrap(contract).token0(),
                        multicallProvider.wrap(contract).token1()
                    ])

                    if (token0 === transferLog.address || token1 === transferLog.address) {
                        let pairAddr = pairsInTx[i];
                        if (!prevUsageInfo.pairs.includes(pairAddr)) prevUsageInfo.pairs.push(pairAddr)
                    }
                } catch (e) {
                    console.error("Unable to get tokens from pair:", e)
                }
            }

            // Set the new info.
            prevUsageInfo.usage += 1;

            // Update.
            tokenUsageMap.set(transferLog.address, prevUsageInfo)
        }))

        // Update last block.
        lastBlockNum = blockNum
    })

    // Start the http server.
    http.createServer(function (req, res) {
        res.writeHead(200, {"Content-Type": "application/json"});
        res.end(JSON.stringify({
            count: sortedUsageMap.size,
            lastUpdate: lastResetDate,
            tokens: Array.from(sortedUsageMap.keys()).map((addr) => {
                let val = sortedUsageMap.get(addr);
                return {
                    "address": addr,
                    "symbol": tokenSymbols.get(addr),
                    "usage": val.usage
                }
            })
        }));
    }).listen(8081);
}

main().catch(e => console.error("Unhandled error: ", e))