package main

import (
	"fmt"
)

func main() {
	fmt.Println("2+3-5= ", sum(2, 3,-5))
	fmt.Println("8+2+4= ", sum(8, 2, 4))
	fmt.Println("7+1= ", sum(7, 1))
}
func sum(xi...int) int {
	su:=0
	for _,v:= range xi{
		su+=v
	}
	return su
}
