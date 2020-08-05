package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//fmt.Println("entered main()")
	s := "ayush shukla"
	ch1 := make(chan string)
	//no need of this channel ch2:=make(chan, string)
	wg.Add(2)
	go rev(ch1)
	ch1 <- s
	fmt.Println(<-ch1)
	go upp(ch1)
	ch1 <- s
	fmt.Println(<-ch1)

	wg.Wait()

}

func rev(ch1 chan string) {
	//fmt.Println("reached rev()")
	str := <-ch1
	st := ""
	for i := range str {
		st = string(str[i]) + st
	}
	ch1 <- st
	//fmt.Println("returned from rev()")
	wg.Done()
}

func upp(ch2 chan string) {
	//fmt.Println("reached upp()")
	st := <-ch2
	ch2 <- strings.ToUpper(st)
	//fmt.Println("returned from upp()")
	wg.Done()

}
