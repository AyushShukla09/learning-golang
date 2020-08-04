package main

import (
  "fmt"
  "strings"
  "os"
)
func words2num(s string) string  {
  var p string
  var b bool
  a:=strings.ToLower(s)
  l:= strings.Split(a, " ")
  for i:=0;i<len(l);i++ {
    switch l[i] {
    case "one":p+="1"
    case "two":p+="2"
    case "three":p+="3"
    case "four":p+="4"
    case "five":p+="5"
    case "six":p+="6"
    case "seven":p+="7"
    case "eight":p+="8"
    case "nine":p+="9"
    case "zero":p+="0"
    default: b=true;break
  }
}
  if b==true{
    return "invalid number"
  }else{
    return p
  }
}
func main(){
  str:=os.Args[1]
  t:=words2num(str)
  fmt.Print(t)
}
