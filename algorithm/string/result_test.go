/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package string

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := "ADOBECODEBANC"
	T := "ABC"
	t.Log(minWindow(s, T))
}

func TestPrint(t *testing.T) {
	chanNum := 4
	chanQueue := make([]chan struct{}, chanNum)
	var result = 0
	exitChan := make(chan struct{})
	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan struct{})
		if i == chanNum-1 {
			go func(i int) {
				chanQueue[i] <- struct{}{}
			}(i)
		}
	}
	for i := 0; i < chanNum; i++ {
		var lastChan, curChan chan struct{}
		if i == 0 {
			lastChan = chanQueue[chanNum-1]
		} else {
			lastChan = chanQueue[i-1]
		}
		curChan = chanQueue[i]
		go func(i byte, lastChan, curChan chan struct{}) {
			for {
				if result > 20 {
					exitChan <- struct{}{}
				}
				<-lastChan
				fmt.Printf("%c\n", i)
				result++
				curChan <- struct{}{}
			}
		}('A'+byte(i), lastChan, curChan)
	}
	<-exitChan
	fmt.Println("done")
}

func TestPrint2(t *testing.T) {
	chanNum := 3
	chanQ := make([]chan int, chanNum)
	exitChan := make(chan struct{})
	for i := 0; i < len(chanQ); i++ {
		chanQ[i] = make(chan int)
		if i == len(chanQ)-1 {
			go func(i2 int) {
				chanQ[i2] <- 1
			}(i)
		}
	}
	for i := 0; i < len(chanQ); i++ {
		var curChan, lastChan chan int
		if i == 0 {
			lastChan = chanQ[chanNum-1]
		} else {
			lastChan = chanQ[i-1]
		}
		curChan = chanQ[i]
		go func(num int, cur, last chan int) {
			for {
				output, ok := <-last
				if ok {
					fmt.Printf("goroutine-%d,print:%d\n", num, output)
					if output < 10 {
						cur <- output + 1
					} else {
						close(cur)
						exitChan <- struct{}{}
					}
				} else {
					_, ok := <-cur
					if ok {
						close(cur)
					}
				}
			}
		}(i, curChan, lastChan)
	}
	<-exitChan
}
