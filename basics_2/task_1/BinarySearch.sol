// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {
    function search(int[] memory arr, int target) public pure returns (int) {
        // 处理空数组情况
        if (arr.length == 0) {
            return -1;
        }
        
        // 初始化左右指针
        int left = 0;
        int right = int(arr.length) - 1;
        
        // 循环查找（当左指针小于等于右指针时）
        while (left <= right) {
            // 计算中间索引（避免(left + right)可能的溢出）
            int mid = left + (right - left) / 2;
            uint midIndex = uint(mid); // 数组索引需要uint类型
            
            // 找到目标值，返回当前索引
            if (arr[midIndex] == target) {
                return mid;
            }
            // 目标值在右半部分，移动左指针
            else if (arr[midIndex] < target) {
                left = mid + 1;
            }
            // 目标值在左半部分，移动右指针
            else {
                right = mid - 1;
            }
        }
        
        // 循环结束仍未找到，返回-1
        return -1;
    }
}