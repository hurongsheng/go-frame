package util

import (
	"fmt"
	"sync"
	"time"
)

/**
利用测试和goroutine并发数控制实现自动驾驶
*/
type sportModule struct {
	speedMode *speedModule
	ch        chan int
	speedList sync.Map
	maxLen    int
}

type waitGroup struct {
	wg        sync.WaitGroup
	len       int
	maxLen    int
	maxLenCh  chan int
	lock      sync.Mutex
	speedMode *speedModule
	sportMode *sportModule
}

//maxLen=0为不限制
func NewWaitGroup() *waitGroup {
	return &waitGroup{
		wg:        sync.WaitGroup{},
		speedMode: &speedModule{},
	}
}

//限制最高并发
func NewLimitedWaitGroup(maxLen int) *waitGroup {
	if maxLen <= 0 {
		return NewWaitGroup()
	}
	sm := &speedModule{}
	wg := &waitGroup{
		wg:        sync.WaitGroup{},
		lock:      sync.Mutex{},
		speedMode: sm,
		maxLen:    maxLen,
		maxLenCh:  make(chan int, maxLen),
	}
	for i := 0; i < wg.maxLen; i++ {
		wg.maxLenCh <- 0
	}
	return wg
}

//限制最高并发
func NewSportWaitGroup(maxLen int) *waitGroup {
	if maxLen <= 0 {
		return NewWaitGroup()
	}
	sm := &speedModule{}
	wg := &waitGroup{
		wg:        sync.WaitGroup{},
		lock:      sync.Mutex{},
		speedMode: sm,
		sportMode: &sportModule{speedMode: sm, maxLen: maxLen, speedList: sync.Map{}},
		maxLen:    maxLen,
		maxLenCh:  make(chan int, maxLen),
	}
	for i := 0; i < wg.maxLen; i++ {
		wg.maxLenCh <- 0
	}
	wg.SetSpeedMode()
	wg.runSportMode()
	return wg
}

func (w *waitGroup) Add(delta int) {
	if w.maxLen > 0 {
		<-w.maxLenCh
	}
	go func() {
		if !w.speedMode.IsOn() {
			return
		}
		w.lock.Lock()
		w.len += delta
		w.lock.Unlock()
		fmt.Printf("len:%d/%d speed: %d \r", w.len, w.maxLen, w.speedMode.Speed())
	}()
	w.wg.Add(delta)
}

func (w *waitGroup) Done() {
	if w.maxLen > 0 {
		w.maxLenCh <- 0
	}
	go func() {
		if !w.speedMode.IsOn() {
			return
		}
		w.lock.Lock()
		w.len--
		w.speedMode.Add(1)
		w.lock.Unlock()
		//测试模式
		fmt.Printf("len:%d/%d speed: %d \r", w.len, w.maxLen, w.speedMode.Speed())
	}()
	w.wg.Done()
}

func (w *waitGroup) Wait() {
	w.wg.Wait()
}

func (w *waitGroup) SetSpeedMode() {
	w.speedMode.On()
}

func (w *waitGroup) runSportMode() {
	t := time.Millisecond * 250
	//t时间测速一次
	go func() {
		for {
			select {
			case <-time.After(t):
				if olds, ok := w.sportMode.speedList.Load(w.maxLen); ok {
					oldsInt := olds.(int)
					if oldsInt > 0 {
						w.sportMode.speedList.Store(w.maxLen, (w.speedMode.Speed()+oldsInt)/2)
					}
				} else {
					w.sportMode.speedList.Store(w.maxLen, w.speedMode.Speed())
				}
				w.speedMode.Reset()
			}
		}
	}()
	//t时间调速一次
	go func() {
		for {
			maxI := 100
			sep := w.maxLen / maxI
			if sep <= 0 {
				sep = 1
			}
			i := 0
			w.changeMaxLen(sep)
			for {
				select {
				case <-time.After(t):
					newLen := w.maxLen
					if i > maxI {
						newLen = w.sportMode.maxSpeedLen()
					} else {
						newLen = (newLen+sep-1)%w.sportMode.maxLen + 1
					}
					w.changeMaxLen(newLen)
					i++
				}
			}
		}
	}()

}
func (w *waitGroup) changeMaxLen(len int) {
	now := w.maxLen
	if len == now {
		return
	}
	if now > len {
		for i := len; i < now; i++ {
			<-w.maxLenCh
		}
	} else {
		for i := now; i < len; i++ {
			w.maxLenCh <- 0
		}
	}
	w.maxLen = len
	fmt.Printf("\n changeMaxLen(%v=>%v) \n", now, len)

}
func (w *sportModule) maxSpeedLen() int {
	max := 0
	maxK := 0
	w.speedList.Range(func(key, value interface{}) bool {
		if value.(int) > max {
			maxK, max = key.(int), value.(int)
		}
		return true
	})
	return maxK
}

/**
测速模块
*/
type speedModule struct {
	start int
	delta int
	on    bool
}

func (s *speedModule) IsOn() bool {
	return s.on
}

func (s *speedModule) On() {
	s.on = true
}

func (s *speedModule) Speed() int {
	return s.delta / (int(time.Now().UnixNano()) - s.start)
}

func (s *speedModule) Add(delta int) int {
	if s.start == 0 {
		s.start = int(time.Now().UnixNano()) - 1
	}
	s.delta += delta * 1000000000
	return s.delta
}

func (s *speedModule) Reset() {
	s.start, s.delta = 0, 0
}
