package main

import (
  "fmt"
  "strconv"
  "os"
)

func num2hex(s string, i uint){
  if s=="-u"{
    H := fmt.Sprintf("%X", i)
    fmt.Print(H)
  }else if s=="-l"{
    h := fmt.Sprintf("%x", i)
    fmt.Print(h)
  }else{
    fmt.Print("Invalid option")
  }
}
func main(){
  if len(os.Args)==2 {
    l,err := strconv.Atoi(os.Args[1])
    if(err==nil){
      num2hex("-l",uint(l))
    } else {
      fmt.Print("Invalid number")
    }
  } else if len(os.Args)==3 {
    l,err := strconv.Atoi(os.Args[2])
    if(err==nil){
      num2hex(os.Args[1],uint(l))
    } else {
      fmt.Print("Invalid number")
    }
  }
}
