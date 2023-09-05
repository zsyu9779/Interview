/*
 * Copyright (C) 2023 Baidu, Inc. All Rights Reserved.
 */
package concurrent

import (
	"context"
	"fmt"
)

type Printer struct {
	ctx       context.Context
	cancel    context.CancelFunc
	gNum      int
	inputChan chan string
}

func NewPrinter(ctx context.Context, gNum int) *Printer {
	ctxC, cancel := context.WithCancel(ctx)
	return &Printer{
		ctx:       ctxC,
		cancel:    cancel,
		gNum:      gNum,
		inputChan: make(chan string, gNum*2),
	}
}

func (p *Printer) Input(text string) {
	p.inputChan <- text
}

func (p *Printer) Start() {
	for i := 0; i < p.gNum; i++ {
		go p.print(p.ctx)
	}

}

func (p *Printer) print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			break
		case text, ok := <-p.inputChan:
			if !ok {
				break
			}
			fmt.Println(text)
		}
	}
}

func (p *Printer) Stop() {
	close(p.inputChan)
	//p.cancel()
}
