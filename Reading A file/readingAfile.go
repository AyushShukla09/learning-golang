package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	content, err := ioutil.ReadFile(`C:\Users\ASHUTOSH\Desktop\self tips.txt`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",content)
}
