//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

/// @dev Supported DEXes
/// Uniswap V2
/// Pancakeswap
/// Biswap
/// SwapV2
/// Nomiswap
/// Babyswap
/// Definix
/// Jetswap
/// BSC
/// StableX
/// Fins
/// Safeswap
/// Warden
/// Elk
/// Panthe
/// Pinkswap
/// DEFISwap
/// SaitaSwap
/// DOOAR
/// LuaSwap
/// Annex
/// FstSwap
/// W3Swap
/// TRUTH

/// @title A contract that connects multiple hooks to one.
/// @author cool guy (@tarik0)
contract FlashloanHooksV1 {
    /// @notice The flashloan hook.
    /// @dev The hook that gets triggered after the funds transfered to contract.
    function onCall(
        address sender,
        uint256 amount0,
        uint256 amount1,
        bytes calldata data
    ) internal virtual {}

    /** The flashswap hooks. */

    function uniswapV2Call(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function pancakeCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function BiswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function swapV2Call(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function nomiswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function babyCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function definixCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function jetswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function BSCswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function stableXCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function FinsCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function safeswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function wardenCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function elkCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function pantherCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function pinkswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function croDefiSwapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function SaitaSwapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function DooarSwapV2Call(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function annexCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function fstswapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function W3swapCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    function TRUTHCall(
        address sender,
        uint amount0,
        uint amount1,
        bytes calldata data
    ) external {
        onCall(sender, amount0, amount1, data);
    }

    // To receive WETH.
    receive() external payable {}
}