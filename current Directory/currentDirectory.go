package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// using the function
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir[strings.LastIndex(mydir, `\`)+1:])
}