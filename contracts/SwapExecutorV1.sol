//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "@uniswap/v2-periphery/contracts/interfaces/IWETH.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./libraries/args/SwapArgsV1.sol";

import "./libraries/SafeMath.sol";

/// @dev Error Codes.
/// SE1: Check your inputs.
/// SE2: Token transfer failed.

/// @title SwapExecutorV1
/// @author cool guy (@tarik0)
/// @notice A helper contract for swaps.
contract SwapExecutorV1 {
    using SafeMath for uint;

    /// @notice Executes a swap between params.Pairs.
    /// @dev Executes a swap between params.Pairs using the parameters.
    function executeSwap(
        SwapParameters calldata params
    ) external {
        // Check inputs.
        require(
            params.Pairs.length + 1 == params.Path.length &&
            params.Path.length == params.AmountsOut.length &&
            params.PairTokens.length == params.Pairs.length,
            "SE1"
        );

        // Send the input token from sender and send it to the pair.
        (bool success,) = address(params.Path[0]).call(
            abi.encodeWithSelector(
                IERC20.transferFrom.selector,
                msg.sender,
                params.Pairs[0],
                params.AmountsOut[0]
            )
        );
        require(success, "SE2");

        // Recursive variables for gas optimization.
        address to;
        uint amount0Out;
        uint amount1Out;

        // Iterate over params.Pairs.
        for (uint i = 0; i < params.Pairs.length; i++) {
            // Calculate "to".
            to = i + 1 < params.Pairs.length ? params.Pairs[i+1] : msg.sender;

            // Calculate amounts.
            (amount0Out, amount1Out) = params.Path[i] == params.PairTokens[i][0] ? (uint(0), params.AmountsOut[i+1]) : (params.AmountsOut[i+1], uint(0));

            // Swap.
            IUniswapV2Pair(params.Pairs[i]).swap(amount0Out, amount1Out, to, new bytes(0));
        }
    }
}

