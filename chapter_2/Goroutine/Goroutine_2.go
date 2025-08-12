package goroutine

import (
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间

type Task func(string) string

type TaskResult struct {
	TaskName      string
	Result        string
	ExecutionTime time.Duration
}

type TaskScheduler struct {
	tasks []struct {
		task  Task
		name  string
		param string
	}
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]struct {
			task  Task
			name  string
			param string
		}, 0),
	}
}

func (s *TaskScheduler) AddTask(name string, task Task, param string) {
	s.tasks = append(s.tasks, struct {
		task  Task
		name  string
		param string
	}{task, name, param})
}

func (s *TaskScheduler) Run() []TaskResult {
	var wg sync.WaitGroup

	resultChan := make(chan TaskResult, len(s.tasks))

	// 为每个任务启动一个协程
	for _, t := range s.tasks {
		wg.Add(1)
		go func(name string, task Task, param string) {
			defer wg.Done()

			// 记录任务开始时间
			startTime := time.Now()

			// 执行任务
			result := task(param)

			// 计算执行时间
			executionTime := time.Since(startTime)

			// 将结果发送到通道
			resultChan <- TaskResult{
				TaskName:      name,
				Result:        result,
				ExecutionTime: executionTime,
			}
		}(t.name, t.task, t.param)
	}

	// 启动一个协程等待所有任务完成后关闭通道
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集所有任务结果
	results := make([]TaskResult, 0, len(s.tasks))
	for result := range resultChan {
		results = append(results, result)
	}
	return results
}
