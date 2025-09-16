// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSortedArrays {
    // 合并两个有序数组（升序），返回新的有序数组
    function mergeArrays(int[] memory arr1, int[] memory arr2) public pure returns (int[] memory) {
        // 处理空数组情况
        if (arr1.length == 0) return arr2;
        if (arr2.length == 0) return arr1;
        
        // 初始化结果数组，长度为两个数组长度之和
        int[] memory result = new int[](arr1.length + arr2.length);
        
        // 定义三个指针：分别指向arr1、arr2和result
        uint i = 0; // arr1的指针
        uint j = 0; // arr2的指针
        uint k = 0; // result的指针
        
        // 双指针遍历，比较元素大小并放入结果数组
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] <= arr2[j]) {
                result[k] = arr1[i];
                i++;
            } else {
                result[k] = arr2[j];
                j++;
            }
            k++;
        }
        
        // 处理arr1的剩余元素
        while (i < arr1.length) {
            result[k] = arr1[i];
            i++;
            k++;
        }
        
        // 处理arr2的剩余元素
        while (j < arr2.length) {
            result[k] = arr2[j];
            j++;
            k++;
        }
        
        return result;
    }

    // 测试函数（可在Remix中直接调用）
    function testMerge() public pure returns (int[] memory) {
        // 测试用例：两个有序数组
        int[] memory arr1 = new int[](3);
        arr1[0] = 1;
        arr1[1] = 3;
        arr1[2] = 5;
        
        int[] memory arr2 = new int[](3);
        arr2[0] = 2;
        arr2[1] = 4;
        arr2[2] = 6;
        
        return mergeArrays(arr1, arr2); // 预期结果: [1,2,3,4,5,6]
    }
}
