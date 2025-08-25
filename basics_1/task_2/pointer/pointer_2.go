package pointer

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func Multiply(s *[]int) {
	for i := range *s {
		(*s)[i] = (*s)[i] * 2
	}
}
