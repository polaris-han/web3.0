package main

import (
	goroutine "chapter_2/Goroutine"
	"chapter_2/lock_mechanism"
	"chapter_2/oop"
	"chapter_2/pointer"
	"fmt"
	"time"
)

// 示例任务1：模拟一个耗时任务
func task1(param string) string {
	time.Sleep(1 * time.Second) // 模拟耗时操作
	return fmt.Sprintf("任务1处理了: %s", param)
}

// 示例任务2：模拟另一个耗时任务
func task2(param string) string {
	time.Sleep(500 * time.Millisecond) // 模拟耗时操作
	return fmt.Sprintf("任务2处理了: %s", param)
}

// 示例任务3：模拟一个更短的任务
func task3(param string) string {
	time.Sleep(200 * time.Millisecond) // 模拟耗时操作
	return fmt.Sprintf("任务3处理了: %s", param)
}

func Goroutine_test() {
	// 创建调度器
	scheduler := goroutine.NewTaskScheduler()

	// 添加任务
	scheduler.AddTask("任务1", task1, "数据A")
	scheduler.AddTask("任务2", task2, "数据B")
	scheduler.AddTask("任务3", task3, "数据C")

	// 运行所有任务并获取结果
	startTime := time.Now()
	results := scheduler.Run()
	totalTime := time.Since(startTime)

	// 打印每个任务的执行结果和时间
	fmt.Println("任务执行结果：")
	for _, result := range results {
		fmt.Printf("%s - 结果: %s, 执行时间: %v\n",
			result.TaskName, result.Result, result.ExecutionTime)
	}

	fmt.Printf("\n所有任务完成，总耗时: %v\n", totalTime)
}

func oop_test() {
	c := oop.Circle{Radius: 2}

	fmt.Printf("circle area %.2f \n", c.Area())
	fmt.Printf("circle perimeter %.2f \n", c.Perimeter())

	r := oop.Rectangle{Width: 3, Height: 5}

	fmt.Printf("Rectangle area %.2f \n", r.Area())
	fmt.Printf("Rectangle perimeter %.2f \n", r.Perimeter())
}

func main() {
	var i int = 10
	fmt.Println(pointer.Add(&i))

	numbers := []int{1, 2, 3, 4, 5}
	pointer.Multiply(&numbers)
	fmt.Println(numbers)

	goroutine.Goroutine_test_1()

	Goroutine_test()

	oop_test()

	var e *oop.Employee = &oop.Employee{
		Person: oop.Person{
			Age:  10,
			Name: "张三",
		},
		EmployeeID: "1234567890",
	}
	e.PrintInfo()

	// channel.TestChannel()
	// channel.TestChannel_2()

	lock_mechanism.TestCounter()
	lock_mechanism.AtomicCounterIncrement()
}
