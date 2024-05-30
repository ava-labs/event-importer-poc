// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

/**
 * @notice Mock contract for mimicking a price feed aggregator.
 * Used in tests to simulate importing price feed updates into other blockchains.
 */
contract MockPriceFeedAggregator {
    address immutable deployer;

    // Latest answer information.
    int256 public currentAnswer;
    uint80 public roundID;
    uint256 public updatedAt;

    event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt);

    constructor() {
        deployer = msg.sender;
    }

    function updateAnswer(int256 _currentAnswer, uint80 _roundID, uint256 _updatedAt) external {
        require(msg.sender == deployer, "Only deployer can update answer");
        currentAnswer = _currentAnswer;
        roundID = _roundID;
        updatedAt = _updatedAt;
        emit AnswerUpdated(_currentAnswer, _roundID, _updatedAt);
    }

    function latestRoundData() public view virtual returns (uint80, int256, uint256, uint256, uint80) {
        return (roundID, currentAnswer, updatedAt, updatedAt, roundID);
    }
}
