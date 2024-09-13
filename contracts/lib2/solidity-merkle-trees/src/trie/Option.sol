pragma solidity ^0.8.24;

import "./Node.sol";

// SPDX-License-Identifier: Apache2

library Option {
    function isSome(OptionalNodeHandle memory val) internal pure returns (bool) {
        return val.op == GenericOption.Some;
    }

    function isSome(NodeHandleOption memory val) internal pure returns (bool) {
        return val.isSome == true;
    }
}
