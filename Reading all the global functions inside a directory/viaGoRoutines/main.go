package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main(){
	var wg sync.WaitGroup		//creating a waitgroup variable
	//s2:=`C:\GOworkSpace\src\revision`    //just in case if my command line dosnt work :)
	s1:=os.Args[1]		//taking input from the command line
	s2:=""
	for j:=0;j<len(s1);j++{		//loop for adding an extra \ so that cammand line input is in readble form
		if string(s1[j])==`\`{		// since '\' is used as escape so adding extra '\' in "C:\GOworkSpace\src\revision"
			s2+=`\`				//to make it in input string readable form by computer
		}
		s2+=string(s1[j])
	}
	//function that reads the names of all the files that within that directory
	//calling the function that pulls name(or path) of all the files and folder within the directory
	var files []string //slice holds all the subdirectory within a file
	err := filepath.Walk(s2, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)				//if you want the name then replace path with info.Name()
		return nil		//appending all the path of files in files[]
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {	//ranging over slice to remove all the folder name and only including files
		if strings.Contains(file, ".") {	//condition for removing folder
			wg.Add(1)				//adding a waitgroup
			go RFunc(file,&wg)		 // calling function to read lines from every file
		}
	}
	wg.Wait() 		//wait till all goroutines are running
}
//function to read lines from every file and
//shortlists all the lines that contain func keyword
func RFunc(sg string, wg *sync.WaitGroup) {
	file, err := os.Open(sg)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	for _, st := range txtlines {
		if strings.Contains(st,"func"){ //shortlisting on basis of "func" keyword
			end:=strings.Index(st, "{") 		//finding the index to take a range of function
			if end!=-1{
				if string(st[5]) == "(" {			//condition for function like "func ()Ayush(){ "
					index := strings.Index(st, ")")
					if st[index+1] >= 65 && st[index+1] <= 90 {		//checking if alphabet is capital after ()
						fmt.Println(st[5:end])
					}
				} else {
					if st[5] >= 65 && st[5] <= 90 {		//condition for checking function like "func Ayush(){ "
						fmt.Println(st[5:end])
					}
				}
			}
		}
	}
	wg.Done()  //closing the waiting waitgroup
}

