/*
 * Copyright (C) 2023 Baidu, Inc. All Rights Reserved.
 */
package concurrent

import "fmt"

func GenerateNatural() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func test() {
	ch := GenerateNatural()
	prime := <-ch
	ch = PrimeFilter(ch, prime)
	for {
		select {
		case a, ok := <-ch:
			if !ok {
				break
			}
			fmt.Println(a)
		}
	}
}
