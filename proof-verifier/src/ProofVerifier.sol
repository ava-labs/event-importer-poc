// SPDX-License-Identifier: Ecosystem
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract ProofVerifier {
    function verifyProof(bytes memory encodedProof, bytes32 leaf, bytes32 root) public view returns (bool) {
        bytes[] memory proof = abi.decode(encodedProof, (bytes[]));

        // Convert proof elements to bytes32
        bytes32[] memory proofElements = new bytes32[](proof.length);
        for (uint256 i = 0; i < proof.length; i++) {
            proofElements[i] = bytes32(proof[i]);
        }

        // Verify the Merkle proof.
        return MerkleProof.verify(proofElements, root, leaf);
    }
}
