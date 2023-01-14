//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

/// @notice The flashloan parameters.
/// @dev Those parameters decodes/encodes inside contract.
struct FlashloanParameters {
        address[] Pairs;
        address[] Path;
        uint[] AmountsOut;
        address[][] PairTokens;
        uint256 BorrowFee;
}
