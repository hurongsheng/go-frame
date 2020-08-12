package util

import (
	"fmt"
	"sync"
)

type waitGroup struct {
	wg       sync.WaitGroup
	len      int
	maxLen   int
	maxLenCh chan int
}

//maxLen=0为不限制
func NewWaitGroup() *waitGroup {
	return &waitGroup{
		wg: sync.WaitGroup{},
	}
}

func NewLimitedWaitGroup(maxLen int) *waitGroup {
	if maxLen <= 0 {
		return NewWaitGroup()
	}
	wg := &waitGroup{
		wg:       sync.WaitGroup{},
		maxLen:   maxLen,
		maxLenCh: make(chan int, maxLen),
	}
	for i := 0; i < wg.maxLen; i++ {
		wg.maxLenCh <- 0
	}
	return wg
}

func (w *waitGroup) RunningSW() int {
	return w.len
}

func (w *waitGroup) Add(delta int) {
	if w.maxLen > 0 {
		fmt.Printf("len:%d/%d\n", w.len, w.maxLen)
		<-w.maxLenCh
	}
	w.len += delta
	w.wg.Add(delta)
}

func (w *waitGroup) Done() {
	w.len--
	if w.maxLen > 0 {
		fmt.Printf("len:%d/%d\n", w.len, w.maxLen)
		w.maxLenCh <- 0
	}
	w.wg.Done()
}

func (w *waitGroup) Wait() {
	w.wg.Wait()
}
