package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 给定10亿个随机数，取其中最大的10个值
func main() {
	rand.Seed(time.Now().UnixNano())
	deep, max := 10000, 10000
	startAt := time.Now().Unix()
	var chDownAt int64
	fmt.Printf("time: %v deep: %v \n", time.Now().Unix(), deep)
	chData := make(chan []int, deep+1)
	chDown := make(chan bool)
	randDataToChan(chData, chDown, deep, max)
	fmt.Printf("time: %v deep: %v \n", time.Now().Unix(), deep)
	flag := true
	flagDown := false
	left := deep * max
	top10Part := make([]int, 0)
	lock := sync.Mutex{}
	for flag || !flagDown {
		select {
		case arr := <-chData:
			if len(arr) == 0 {
				flagDown = true
			}
			fmt.Printf("left(%v,%v)  %+v\n", left, len(arr), arr)
			go func(arr []int) {
				top10Part = append(top10Part, getTop10Array(arr)...)
				lock.Lock()
				left = left - max + 10
				if len(top10Part) >= max {
					chData <- top10Part[:max]
					top10Part = top10Part[max:]
				}
				lock.Unlock()
			}(arr)

		case flagDown = <-chDown:
			chDownAt = time.Now().Unix()
		default:
			if flagDown {
				flag = false
			}
		}
	}
	fmt.Printf("left(%v,%v)  %+v\n", left, len(top10Part), top10Part)
	top10Part = getTop10Array(top10Part)
	fmt.Printf("time: %v last top10Part %+v\n", time.Now().Unix(), top10Part)
	fmt.Printf("start at: %v,ch down at %v,end at %+v\n", startAt, chDownAt, time.Now().Unix())

}

func randDataToChan(chData chan []int, chDown chan bool, deep int, max int) {
	sw := sync.WaitGroup{}
	sw.Add(deep)
	for j := 0; j < deep; j++ {
		go func(j int) {
			defer sw.Done()
			chData <- randArray2(max, max*deep)
		}(j)
	}
	go func() {
		sw.Wait()
		chDown <- true
	}()
	fmt.Printf("time: %v rand down", time.Now().Unix())
}

func getTop10Array(valueArr []int) []int {
	l := len(valueArr)
	if l <= 10 {
		return valueArr
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < l-i-1; j++ {
			if valueArr[j] > valueArr[j+1] {
				valueArr[j+1], valueArr[j] = valueArr[j], valueArr[j+1]
			}
		}
	}
	return valueArr[l-10:]
}

func randArray2(num int, max int) []int {
	arr := make([]int, 0, num)
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(max))
	}
	return arr
}
