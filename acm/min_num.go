package main

import (
	"fmt"
	. "frame/util"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("need arr")
		return
	}
	req := os.Args[1]
	arr := strings.Split(req, ",")
	var minN *int
	for _, numStr := range arr {
		n := StrToInt(numStr)
		if minN == nil {
			minN = &n
			continue
		}
		if *minN > n {
			minN = &n
		}
	}
	fmt.Printf("min num %v\n", *minN)
}
