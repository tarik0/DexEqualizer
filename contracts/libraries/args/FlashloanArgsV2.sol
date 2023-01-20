//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

/// @notice The flashloan parameters.
/// @dev Those parameters decodes/encodes inside contract.
struct FlashloanParameters {
          // Pairs.
          address[] Pairs;
          uint256[][] Reserves;
          address[][] PairTokens;

          // Path.
          address[] Path;
          uint[] AmountsOut;

          // Chi.
          address GasToken;
          uint256 GasTokenAmount;

          // Other parameters.
          uint256 PoolDebt;
          bool RevertOnReserveChange;
}
