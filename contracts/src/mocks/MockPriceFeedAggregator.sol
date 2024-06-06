// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Mock contract for mimicking a price feed aggregator.
 * Used in tests to simulate importing price feed updates into other blockchains.
 */
contract MockPriceFeedAggregator {
    address public immutable deployer;

    // Latest answer information.
    int256 public currentAnswer;
    uint80 public roundID;
    uint256 public updatedAt;

    event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt);

    constructor() {
        deployer = msg.sender;
    }

    function updateAnswer(int256 currentAnswer_, uint80 roundID_, uint256 updatedAt_) external {
        require(msg.sender == deployer, "Only deployer can update answer");
        currentAnswer = currentAnswer_;
        roundID = roundID_;
        updatedAt = updatedAt_;
        emit AnswerUpdated(currentAnswer_, roundID_, updatedAt_);
    }

    function latestRoundData() public view virtual returns (uint80, int256, uint256, uint256, uint80) {
        return (roundID, currentAnswer, updatedAt, updatedAt, roundID);
    }
}
