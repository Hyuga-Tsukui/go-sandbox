package worker

import (
	"context"
	"fmt"
	"log"
)

type Task struct {
	ID    int
	Argas any
}

type EmailTask struct {
	To      string
	Subject string
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

	switch task.Argas.(type) {
	case EmailTask:
		log.Printf("send email: %v", task.Argas.(EmailTask))
	default:
		log.Printf("exectute task: %v", task)
	}

	// シミュレーターとして3秒待つ
	// time.Sleep(3 * time.Second)
	//ランダムでpanicを発生させる
	// if rand.Intn(2) == 0 {
	// 	panic("panic")
	// }
}
