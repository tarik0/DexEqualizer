import "@uniswap/v2-periphery/contracts/interfaces/IWETH.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./libraries/hooks/FlashloanHooksV1.sol";
import "./libraries/args/FlashloanArgsV2.sol";

import "./libraries/SafeMath.sol";

/// @dev ChiToken interface.
interface IChiToken {
    function freeFromUpTo(address from, uint256 value) external;
}

/// @title FlashloanExecutorV2
/// @author cool guy (@tarik0)
/// @notice A helper contract for flashloan swaps.
contract FlashloanExecutorV2 is FlashloanHooksV1 {
    using SafeMath for uint;

    /// @notice Starts the whole flashloan process.
    /// @dev Triggers the flashswap.
    function executeFlashloan(
        FlashloanParameters calldata params
    ) external {
        // Check inputs.
        require(
            params.Pairs.length + 1 == params.Path.length &&
            params.Path.length == params.AmountsOut.length &&
            params.PairTokens.length == params.Pairs.length,
            "FE1"
        );

        // Check reserves.
        for (uint i = 0; i < params.Pairs.length; i++) {
            (uint r0, uint r1,) = IUniswapV2Pair(params.Pairs[i]).getReserves();
            if (r0 != params.Reserves[i][0] || r1 != params.Reserves[i][1]) {
                if (params.RevertOnReserveChange) require(false, "FE2");
                return;
            }
        }

        // Select amounts.
        (
        uint amount0Out,
        uint amount1Out
        ) = params.Path[0] == params.PairTokens[0][0] ? (uint(0), params.AmountsOut[1]) : (params.AmountsOut[1], uint(0));

        // Trigger the flashloan.
        IUniswapV2Pair(params.Pairs[0]).swap(
            amount0Out,
            amount1Out,
            address(this),
            abi.encode(params)
        );
    }

    /// @notice A global hook to capture all flashswap hooks.
    /// @dev It's from FlashloanHooks.
    function onCall(
        address,
        uint256,
        uint256,
        bytes calldata data
    ) internal override {
        // Check data.
        require(data.length > 0, "FE3");

        // Decode parameters.
        FlashloanParameters memory params = abi.decode(data, (FlashloanParameters));

        // Check if someone else triggered.
        require(msg.sender == params.Pairs[0], "FE4");

        // Recursive variables for gas optimization.
        address to;
        uint amount0Out;
        uint amount1Out;

        // Send the input token from sender and send it to the pair.
        (bool success,) = address(params.Path[1]).call(
            abi.encodeWithSelector(
                IERC20.transfer.selector,
                params.Pairs[1],
                params.AmountsOut[1]
            )
        );
        require(success, "FE5");

        // Iterate over pairs.
        for (uint i = 1; i < params.Pairs.length; i++) {
            // Calculate "to".
            to = i + 1 < params.Pairs.length ? params.Pairs[i+1] : address(this);

            // Calculate amounts.
            (
            amount0Out,
            amount1Out
            ) = params.Path[i] == params.PairTokens[i][0] ? (uint(0), params.AmountsOut[i+1]) : (params.AmountsOut[i+1], uint(0));

            // Swap.
            IUniswapV2Pair(params.Pairs[i]).swap(amount0Out, amount1Out, to, new bytes(0));
        }

        // Pay the debt.
        (success,) = address(params.Path[0]).call(
            abi.encodeWithSelector(
                IERC20.transfer.selector,
                params.Pairs[0],
                params.PoolDebt
            )
        );
        require(success, "FE6");

        // Return profit.
        (success,) = address(params.Path[0]).call(
            abi.encodeWithSelector(
                IERC20.transfer.selector,
                tx.origin,
                params.AmountsOut[params.AmountsOut.length-1].sub(params.PoolDebt)
            )
        );
        require(success, "FE7");

        // Stop here if you don't want to use gas token.
        if (params.GasTokenAmount == 0) return;

        // Burn gas token.
        (success,) = address(params.GasToken).call(
            abi.encodeWithSelector(
                IChiToken.freeFromUpTo.selector,
                tx.origin,
                params.GasTokenAmount
            )
        );
        require(success, "FE8");
    }
}