// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegerToRoman {
    // 定义罗马数字的数值与对应符号（从大到小排列）
    uint256[] private values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
    string[] private symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];

    // 将整数转换为罗马数字（支持 1-3999）
    function intToRoman(uint256 num) public view returns (string memory) {
        // 检查输入范围（罗马数字通常不表示0，且超过3999需要特殊符号）
        require(num > 0 && num <= 3999, "Number must be between 1 and 3999");

        string memory result = "";

        // 贪心算法：从最大的数值开始匹配
        for (uint256 i = 0; i < values.length; i++) {
            // 当当前数值小于等于剩余数字时，拼接符号并减去对应值
            while (num >= values[i]) {
                result = string(abi.encodePacked(result, symbols[i]));
                num -= values[i];
            }
            
            // 数字减为0时提前退出循环
            if (num == 0) {
                break;
            }
        }

        return result;
    }

    // 测试函数（可在Remix中直接调用查看结果）
    function testConversion(uint256 num) public view returns (string memory) {
        return intToRoman(num);
    }
}
