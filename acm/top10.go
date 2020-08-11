package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	deep, max := 10000, 10000
	fmt.Printf("time: %v deep: %v \n", time.Now().Unix(), deep)
	arr := randData(deep, max)
	fmt.Printf("time: %v deep: %v \n", time.Now().Unix(), deep)
	top10Part := sync.Map{}
	for {
		fmt.Printf("a time: %v deep: %v \n", time.Now().Unix(), deep)
		top10Part = getTop10(arr)
		fmt.Printf("b time: %v deep: %v \n", time.Now().Unix(), deep)
		arr, deep = mergeTop10(top10Part, max)
		fmt.Printf("c time: %v deep: %v \n", time.Now().Unix(), deep)
		if deep == 1 {
			top10Part = getTop10(arr)
			break
		}
	}
	fmt.Printf("time: %v deep: %v \n", time.Now().Unix(), deep)
	top10Part.Range(func(key, value interface{}) bool {
		fmt.Printf("max top 10 %+v\n", value)
		return true
	})
}

func randData(deep int, max int) sync.Map {
	sw := sync.WaitGroup{}
	arr := sync.Map{}
	sw.Add(deep)
	for j := 0; j < deep; j++ {
		go func(j int) {
			arr.Store(j, randArray(max, max*100))
			sw.Done()
		}(j)
	}
	sw.Wait()
	return arr
}

func getTop10(arr sync.Map) sync.Map {
	top10Part := sync.Map{}
	sw := sync.WaitGroup{}
	arr.Range(func(key, value interface{}) bool {
		sw.Add(1)
		go func(key int) {
			defer sw.Done()
			valueArr := value.([]int)
			l := len(valueArr)
			if l <= 10 {
				top10Part.Store(key, valueArr)
				return
			}
			for i := 0; i < 10; i++ {
				for j := 0; j < l-i-1; j++ {
					if valueArr[j] > valueArr[j+1] {
						valueArr[j+1], valueArr[j] = valueArr[j], valueArr[j+1]
					}
				}
			}
			top10Part.Store(key, valueArr[l-10:])
		}(key.(int))
		return true
	})
	sw.Wait()
	return top10Part
}
func mergeTop10(top10Part sync.Map, max int) (sync.Map, int) {
	top10Arr := make(map[int][]int, 0)
	top10PartSyncMap := sync.Map{}
	merge := max / 10
	mergeFlag := 0
	top10Part.Range(func(key, value interface{}) bool {
		top10Arr[key.(int)/merge] = append(top10Arr[key.(int)/merge], value.([]int)...)
		mergeFlag = key.(int) / merge
		if (key.(int) / merge) == (merge - 1) {
			top10PartSyncMap.Store(key.(int)/merge, top10Arr[key.(int)/merge])
		}
		return true
	})
	top10PartSyncMap.Store(mergeFlag, top10Arr[mergeFlag])
	return top10PartSyncMap, len(top10Arr)
}

func randArray(num int, max int) []int {
	arr := make([]int, 0, num)
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(max))
	}
	return arr
}
