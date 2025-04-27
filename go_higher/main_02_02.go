package main

import (
	"fmt"
	"time"
)

// ✅Goroutine
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func main() {
	// 定义一些示例任务
	tasks := []Task{
		func() {
			fmt.Println("Task 1 started")
			time.Sleep(2 * time.Second)
			fmt.Println("Task 1 finished")
		},
		func() {
			fmt.Println("Task 2 started")
			time.Sleep(1 * time.Second)
			fmt.Println("Task 2 finished")
		},
		func() {
			fmt.Println("Task 3 started")
			time.Sleep(3 * time.Second)
			fmt.Println("Task 3 finished")
		},
	}

	// 创建任务调度器并运行任务
	scheduler := NewTaskScheduler(tasks)
	results := scheduler.Run()

	// 打印每个任务的执行时间和结果
	for _, result := range results {
		fmt.Printf("Task %d took %v to complete\n", result.Index+1, result.Duration)
	}

}

type Task func()

type TaskResult struct {
	Index    int
	Duration time.Duration
	Result   interface{}
}

type TaskScheduler struct {
	tasks []Task
}

func NewTaskScheduler(tasks []Task) *TaskScheduler {
	return &TaskScheduler{
		tasks: tasks,
	}
}
func (s *TaskScheduler) Run() []TaskResult {
	results := make(chan TaskResult, len(s.tasks))
	defer close(results)

	for i, task := range s.tasks {
		go func(index int, t Task) {
			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)
			results <- TaskResult{
				Index:    index,
				Duration: duration,
				Result:   nil, // 如果任务有返回值，可以在这里收集
			}
		}(i, task)
	}

	var taskResults []TaskResult
	for i := 0; i < len(s.tasks); i++ {
		result := <-results
		taskResults = append(taskResults, result)
	}

	return taskResults
}
