// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {MerkleProof} from "@openzeppelin/contracts@4.8.1//utils/cryptography/MerkleProof.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RlpReader.sol";

struct EVMLog {
    address loggerAddress;
    bytes32[] topics;
    bytes data;
}

struct EVMReceipt {
    bool status;
    uint64 cululativeGasUsed;
    bytes bloom;
    EVMLog[] logs;
}

contract ProofVerifier {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;

    function verifyProof(bytes memory encodedProof, bytes32 leaf, bytes32 root) public pure returns (bool) {
        bytes[] memory proof = abi.decode(encodedProof, (bytes[]));

        // Convert proof elements to bytes32
        bytes32[] memory proofElements = new bytes32[](proof.length);
        for (uint256 i = 0; i < proof.length; i++) {
            proofElements[i] = bytes32(proof[i]);
        }

        // Verify the Merkle proof.
        return MerkleProof.verify(proofElements, root, leaf);
    }

    function verifyTrieProof(bytes32 root, bytes[] memory proof, bytes memory key) public pure returns (bytes memory) {
        // Put the single key as an array
        bytes[] memory keys = new bytes[](1);
        keys[0] = key;

        StorageValue[] memory results = MerklePatricia.VerifyEthereumProof(root, proof, keys);
        return results[0].value;
    }

    function getFirstReceipt(bytes32 receiptRoot, bytes[] memory proof) public pure returns (EVMReceipt memory) {
        bytes memory encodedReceipt = verifyTrieProof(receiptRoot, proof, hex"80");
        return decodeReceipt(encodedReceipt);
    }

    function decodeReceipt(bytes memory encodedReceipt) public pure returns (EVMReceipt memory) {
        RLPReader.RLPItem[] memory receipt = encodedReceipt.toRlpItem().toList();
        require(receipt.length == 4, "Invalid number of RLP elements in receipt");
        EVMReceipt memory evmReceipt;
        evmReceipt.status = receipt[0].toUint() == 1;
        evmReceipt.cululativeGasUsed = uint64(receipt[1].toUint());
        evmReceipt.bloom = receipt[2].toBytes();
        RLPReader.RLPItem[] memory logs = receipt[3].toList();
        evmReceipt.logs = new EVMLog[](logs.length);
        for (uint256 i = 0; i < logs.length; i++) {
            RLPReader.RLPItem[] memory log = logs[i].toList();
            require(log.length == 3, "Invalid number of RLP elements in log");
            evmReceipt.logs[i].loggerAddress = log[0].toAddress();
            RLPReader.RLPItem[] memory topics = log[1].toList();
            evmReceipt.logs[i].topics = new bytes32[](topics.length);
            for (uint256 j = 0; j < topics.length; j++) {
                evmReceipt.logs[i].topics[j] = bytes32(topics[j].toBytes());
            }
            evmReceipt.logs[i].data = log[2].toBytes();
        }
        return evmReceipt;
    }
}
