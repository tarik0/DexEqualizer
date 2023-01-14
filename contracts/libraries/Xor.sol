pragma solidity ^0.8.0;

/// @title XOR
/// @dev An XOR library to encrypt/decrypt
library XOR {
    /// @dev crypt encrypts/decrypts the data by the given key.
    function crypt (bytes memory data, bytes memory key) internal pure returns (bytes memory result) {
        // Store data length on stack for later use
        uint256 length = data.length;

        assembly {
        // Set result to free memory pointer
            result := mload (0x40)
        // Increase free memory pointer by lenght + 32
            mstore (0x40, add (add (result, length), 32))
        // Set result length
            mstore (result, length)
        }

        // Iterate over the data stepping by 32 bytes
        for (uint i = 0; i < length; i += 32) {
            // Generate hash of the key and offset
            bytes32 hash = keccak256 (abi.encodePacked (key, i));

            bytes32 chunk;
            assembly {
            // Read 32-bytes data chunk
                chunk := mload (add (data, add (i, 32)))
            }
            // XOR the chunk with hash
            chunk ^= hash;
            assembly {
            // Write 32-byte encrypted chunk
                mstore (add (result, add (i, 32)), chunk)
            }
        }
    }

}