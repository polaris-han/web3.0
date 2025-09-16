// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    mapping(string => uint256) public voteCount;

    function vote(string memory candidate) public {
        voteCount[candidate] += 1;
    }

    function getVotes(string memory candidate) public view returns (uint256) {
        return voteCount[candidate];
    }

    function resetVotes(string memory candidate) public {
        voteCount[candidate] = 0;
    }

}