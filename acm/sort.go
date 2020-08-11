package main

import (
	"fmt"
	. "frame/util"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("need sort type")
		return
	}
	if len(os.Args) <= 2 {
		fmt.Println("need arr")
		return
	}
	sortType := os.Args[1]
	var arr []int
	if os.Args[2] == "rand" {
		arr = randArr(10, 30)
	} else {
		arr = StrToInts(os.Args[2], ",")
	}
	fmt.Printf("%v %+v\n", sortType, arr)
	switch sortType {
	case "quick":
		arr = QuickSort(arr)
	case "bubble":
		arr = BubbleSort(arr)
	}
	fmt.Printf("%v %+v\n", sortType, arr)

}

func randArr(num int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 0, num)
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(max))
	}
	return arr
}

//O(N*logN)
func QuickSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[i] > mid {
			fmt.Printf("arr %+v,head %v,tail %v mid %v  switch(%v,%v) \n", values, head, tail, mid, i, tail)
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			fmt.Printf("arr %+v,head %v,tail %v mid %v  switch(%v,%v) \n", values, head, tail, mid, i, head)
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	fmt.Printf("arr %+v,head %v,tail %v mid %v \n", values, head, tail, mid)

	QuickSort(values[:head])
	QuickSort(values[head+1:])
	return values
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		s := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				s = true
			}
		}
		if !s {
			break
		}
		fmt.Printf("arr: %+v\n", arr)
	}
	return arr
}
