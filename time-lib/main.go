package main

import (
	"fmt"
	"time"
)

func main() {
	ct := time.Now()
	// 計測したい処理
	// シミュレーションとして1秒待つ
	time.Sleep(1 * time.Second)
	fmt.Printf("processingTime: %v\n", time.Since(ct))
}
