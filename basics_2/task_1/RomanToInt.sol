// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RomanToInt {

    /* 单个字符 → 数值 */
    function _valueOf(bytes1 c) private pure returns (uint16) {
        if (c == "I") return 1;
        if (c == "V") return 5;
        if (c == "X") return 10;
        if (c == "L") return 50;
        if (c == "C") return 100;
        if (c == "D") return 500;
        if (c == "M") return 1000;
        return 0; // 非法
    }

    function romanToInt(string memory s) public pure returns (uint16) {
        bytes memory param = bytes(s);

        uint16 acc;
        uint16 prev;

        for(uint256 i = param.length; i > 0; ) {
            --i;
            uint16 cur = _valueOf(param[i]);

            if (cur == 0) revert InvalidRoman(); // 非法字符

            acc = cur < prev ? acc - cur : acc + cur;
            prev = cur;           
        }

        return acc;
    }

    error InvalidRoman();
}