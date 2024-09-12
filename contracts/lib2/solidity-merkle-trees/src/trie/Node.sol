pragma solidity 0.8.18;

// SPDX-License-Identifier: Apache2

import "./NibbleSlice.sol";
import "./Bytes.sol";

/// This is an enum for the different node types.
struct NodeKind {
    bool isEmpty;
    bool isLeaf;
    bool isHashedLeaf;
    bool isNibbledValueBranch;
    bool isNibbledHashedValueBranch;
    bool isNibbledBranch;
    bool isExtension;
    bool isBranch;
    uint256 nibbleSize;
    ByteSlice data;
}

struct NodeHandle {
    bool isHash;
    bytes32 hash;
    bool isInline;
    bytes inLine;
}

struct Extension {
    NibbleSlice key;
    NodeHandle node;
}

struct Branch {
    OptionalNodeHandle value;
    OptionalNodeHandle[16] children;
}

struct NibbledBranch {
    NibbleSlice key;
    OptionalNodeHandle value;
    OptionalNodeHandle[16] children;
}

struct ValueOption {
    bool isSome;
    bytes value;
}

struct NodeHandleOption {
    bool isSome;
    NodeHandle value;
}

struct Leaf {
    NibbleSlice key;
    NodeHandle value;
}

struct TrieNode {
    bytes32 hash;
    bytes node;
}

enum GenericOption {
    None,
    Some
}

struct OptionalNodeHandle {
    GenericOption op;
    uint256 ptr;
}

library OptionalNodeHandleLib {
    using OptionalNodeHandleLib for OptionalNodeHandle;

    // function isNone(OptionalNodeHandle memory o) internal pure returns (bool) {
    //     return o.op == GenericOption.None;
    // }

    // function isSome(OptionalNodeHandle memory o) internal pure returns (bool) {
    //     return o.op == GenericOption.Some;
    // }

    function unwrap(OptionalNodeHandle memory o) internal pure returns (NodeHandle memory) {
        // if (o.isSome()) {
        if (o.op == GenericOption.Some) {
            return o.unwrapUnchecked();
        }
        revert("panic!");
    }

    function unwrapUnchecked(OptionalNodeHandle memory o) internal pure returns (NodeHandle memory nodeHandle) {
        uint256 ptr = o.ptr;
        assembly {
            nodeHandle := ptr
        }
        return nodeHandle;
    }

    function none() internal pure returns (OptionalNodeHandle memory) {
        return OptionalNodeHandle({
            op: GenericOption.None,
            ptr: 0
        });
    }

    function some(NodeHandle memory nodeHandle) internal pure returns (OptionalNodeHandle memory) {
        uint256 ptr;
        assembly {
            ptr := nodeHandle
        }
        return OptionalNodeHandle({
            op: GenericOption.Some,
            ptr: ptr
        });
    }
}
