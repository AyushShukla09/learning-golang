package main

import (
	"bufio"
	"log"
	"os"
)
var str []string
func main() {
	file, err := os.Open(`C:\GOworkSpace\src\revision\new 3.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str=append(str,scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//till here we have a slice which contains a line by line of a file
	for _,v := range s[4]{
		if v==32{
			println()
			continue
		}
		print(string(v))

	}


}
