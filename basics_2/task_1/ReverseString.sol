// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ReverseString {
    function reverse(string memory param) public pure returns (string memory) {
        bytes memory paramArr = bytes(param);
        bytes memory temp = new bytes(paramArr.length);

        for (uint256 i = 0; i < paramArr.length; i++) {
            temp[i] = paramArr[paramArr.length - i - 1];
        }
        return string(temp);

    }
}