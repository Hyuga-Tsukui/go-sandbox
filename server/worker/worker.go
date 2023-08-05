package worker

import (
	"context"
	"fmt"
	"log"
)

type Task struct {
	ID int
}

func RunWorker(ctx context.Context, buf int) (queue chan *Task) {
	queue = make(chan *Task, buf)
	errChan := make(chan error)

	go func() {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case task := <-queue:
				go executeTask(task, errChan)
			case err := <-errChan:
				log.Printf("error: %v", err)
			}
		}
		log.Println("worker closed")
		close(queue)
		close(errChan)
	}()
	return queue
}

func executeTask(task *Task, errChan chan error) {
	defer func() {
		if r := recover(); r != nil {
			errChan <- fmt.Errorf("panic: %v", r)
		}
	}()

	// シミュレーターとして3秒待つ
	// time.Sleep(3 * time.Second)
	//ランダムでpanicを発生させる
	// if rand.Intn(2) == 0 {
	// 	panic("panic")
	// }
	log.Printf("exectute task: %v", task)
}
