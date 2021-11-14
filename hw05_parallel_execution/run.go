package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	var ret error
	errCount := 0
	numTasks := len(tasks)
	ch := make(chan int, numTasks)
	errChan := make(chan int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				num, ok1 := <-ch
				if !ok1 {
					break
				}
				err := tasks[num]()
				mu.Lock()
				if err != nil {
					errCount++
				}
				mu.Unlock()
				errChan <- errCount
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numTasks; i++ {
			ch <- i
			if <-errChan >= m {
				ret = ErrErrorsLimitExceeded
				break
			}
		}
		close(ch)
	}()
	wg.Wait()
	return ret
}
