/*
 * Copyright (C) 2023 Baidu, Inc. All Rights Reserved.
 */
package concurrent

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	printer := NewPrinter(context.Background(), 3) // 创建一个打印机对象，创建3个worker

	// 输入要打印的内容
	printer.Input("Hello")
	printer.Input("World")
	// 调用Start方法开始打印
	go printer.Start()

	// 持续输入内容
	go func() {
		for i := 0; i < 5; i++ {
			printer.Input(fmt.Sprintf("Input %d", i+1))
		}
	}()

	// 等待一段时间后调用Stop方法结束打印
	// 也可以根据实际需求进行控制，例如通过用户输入来判断何时结束打印
	<-time.After(5 * time.Second)
	printer.Stop()
}

func TestFor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	testCh := make(chan int, 10)
	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 100; i++ {
			testCh <- i
		}
		//close(testCh)
		cancel()
	}()
	for i := 0; i < 3; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				case a, ok := <-testCh:
					if !ok {
						break
					}
					fmt.Println(a)
				}
			}
		}()
	}
	time.Sleep(5 * time.Second)
}
