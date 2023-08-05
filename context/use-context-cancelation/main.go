// REF: https://zenn.dev/hsaki/books/golang-context/viewer/done

package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}
		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

// Contextを利用して、ゴルーチンへの終了伝達を行うパターンの実装.
// Cancel可能なContext（WithCancel）を作成し、Cancelをよぶことでゴルーチンへ終了を伝える.
func main() {

	ctx, cancel := context.WithCancel(context.Background())

	gen := generator(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	// メインゴルーチンが終了する前にgeneratorが終了するのを待つ.
	// メインゴルーチンが終了すると、generatorが強制終了する.
	wg.Wait()
}
