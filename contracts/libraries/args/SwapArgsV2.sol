//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

/// @notice The swap parameters.
/// @dev Those parameters decodes/encodes inside contract.
struct SwapParameters {
        // Pairs.
        address[] Pairs;
        uint256[][] Reserves;
        address[][] PairTokens;

        // Path.
        address[] Path;
        uint[] AmountsOut;

        // Other.
        bool RevertOnReserveChange;
        address GasToken;
        bool UseGasToken;
}
