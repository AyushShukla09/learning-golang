package main

import (
  "fmt"   //for formatting
  "os"  //for user entry
)
// A function to convert number to words and "Not a number" incase of invalid entry
func num2words(s string) string  {
  var b bool
  var p string
  for i:=0;i<len(s);i++ {
    switch string(s[i]) {
    case "1":p+="one "
    case "2":p+="two "
    case "3":p+="three "
    case "4":p+="four "
    case "5":p+="five "
    case "6":p+="six "
    case "7":p+="seven "
    case "8":p+="eight "
    case "9":p+="nine "
    case "0":p+="zero "
    case " ":p+="\n"
    default: b=true;break
    }
  }
  if b==false{
    return p
  }else{
    return ("Not a number")
  }
}
//excetion starts from here
func main(){
  for i:=1;i<len(os.Args);i++{
    st:=os.Args[i]
    t:=num2words(st)
    fmt.Println(t)
  }
  }
