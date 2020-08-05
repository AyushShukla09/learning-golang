package main

import (
	"fmt"
	"time"
)

func findmin(a []int) {
	var min int=a[0]
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
	}
	fmt.Println(a, "--->", min)
}
func main() {

	var x = []int{5, 8, 2, 12, 1}
	var y = []int{15, -18, -12, 1, 3}
	go findmin(x)
	go findmin(y)
	time.Sleep(time.Duration(1) * time.Millisecond)
}
