//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

/// @notice The swap parameters.
/// @dev Those parameters gets used when swap.
struct SwapParameters {
    address[] Pairs;
    address[] Path;
    uint256[] AmountsOut;
    address[][] PairTokens;
}
