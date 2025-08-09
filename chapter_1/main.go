package main

import (
	"fmt"
	"math"
	"strconv"
)

func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for numsIndex := range nums {
		numMap[nums[numsIndex]]++
	}
	for key, value := range numMap {
		if value == 1 {
			return key
		}
	}
	return -1 // or another appropriate default value
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}

func isValid(s string) bool {
	stack := []rune{}
	// 键 为结束括号，值 为开始括号
	matching := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, char := range s {
		// 当前括号是否存在于matching的 键
		key, exists := matching[char]
		if exists {
			// 存在 当前括号为结束括号，若栈（stack）的长度为0或 栈顶的元素 不是 开始括号，则括号不完整 返回 false
			if len(stack) == 0 || stack[len(stack)-1] != key {
				return false
			} else {
				// 否则，即栈顶元素为 开始括号，将元素弹出
				// 此操作 左闭右开 区间
				stack = stack[:len(stack)-1]
			}
		} else {
			// 不存在，当前括号为开始括号，栈顶 放入当前括号
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	var minLength int = len(strs[0])
	for _, str := range strs {
		if minLength > len(str) {
			minLength = len(str)
		}
	}

	var prefix string
	for i := 1; i <= minLength; i++ {
		var tmpPrefix []rune = ([]rune(strs[0]))[:i]
		var isCommonPrefix bool = true
		for _, str := range strs {
			if "" == str {
				return ""
			}
			tempRune := ([]rune(str))[:i]
			for j := range tmpPrefix {
				if tmpPrefix[j] != tempRune[j] {
					isCommonPrefix = false
				}
			}
		}
		if isCommonPrefix {
			prefix = string(tmpPrefix)
		}
	}
	return prefix
}

func plusOne(digits []int) []int {
	var num int = 0
	for i, digit := range digits {
		var j int = len(digits) - 1 - i
		num += int(digit * int(math.Pow(10, float64(j))))
	}
	num += 1
	numStr := strconv.Itoa(num)
	var result []int
	for _, char := range numStr {
		digit := int(char - '0')
		result = append(result, digit)
	}
	return result
}

func removeDuplicates(nums []int) int {
	var pointer2 []int
	var numMap map[int]int = make(map[int]int)
	for index, value := range nums {
		num, exist := numMap[value]
		// fmt.Println(index, value)
		if !exist {
			numMap[value] = num
			// 第一个不重复数字的下标
			pointer2 = append(pointer2, index)
		}
	}
	if len(pointer2) > 0 {
		for i := 0; i < len(pointer2); i++ {
			nums[i] = nums[pointer2[i]]
		}
	}
	return len(pointer2)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	var result [][]int = [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := result[len(result)-1]

		if current[0] <= last[1] {
			result[len(result)-1][1] = max(last[1], current[1])
		} else {
			result = append(result, current)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// nums := []int{2, 2, 1}
	// fmt.Println(singleNumber(nums))

	// x := 121
	// fmt.Println(isPalindrome(x))
	// x = -121
	// fmt.Println(isPalindrome(x))
	// x = 10
	// fmt.Println(isPalindrome(x))

	// fmt.Println(isValid("()"))
	// fmt.Println(isValid("()[]{}"))
	// fmt.Println(isValid("(]"))
	// fmt.Println(isValid("([])"))
	// fmt.Println(isValid("([)]"))

	// 	var strs []string = []string{"flower", "flow", "flight"}
	// 	fmt.Println(longestCommonPrefix(strs))
	// 	strs = []string{"dog", "racecar", "car"}
	// 	fmt.Println(longestCommonPrefix(strs))
	// 	strs = []string{"flower", "flow", "flight", ""}
	// 	fmt.Println(longestCommonPrefix(strs))

	// fmt.Println(plusOne([]int{1, 2, 3}))
	// fmt.Println(plusOne([]int{4, 3, 2, 1}))

	// nums := []int{1, 1, 2}
	// fmt.Println("有效长度：", removeDuplicates(nums), " 新数组：", nums)
	// nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// fmt.Println("有效长度：", removeDuplicates(nums), " 新数组：", nums)

	var intervals [][]int = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))
}
