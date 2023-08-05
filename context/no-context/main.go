package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case out <- num:
			}
		}
		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

// Contextを利用せずに、ゴルーチンへの終了伝達を行うパターンの実装.
// 終了用のチャネルを作成し、ゴルーチンの終了時にチャネルを閉じる.
func main() {
	// チャネルが閉じられたかどうかを確認するためのチャネルを作成する.
	// このようなチャネルはメモリの効率から空構造体を送る.
	done := make(chan struct{})
	gen := generator(done, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	close(done)

	// メインゴルーチンが終了する前にgeneratorが終了するのを待つ.
	// メインゴルーチンが終了すると、generatorが強制終了する.
	wg.Wait()
}
