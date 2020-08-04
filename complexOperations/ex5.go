package main

import (
  "fmt"
  )

func main(){
  var a complex128=3-5i
  var b complex128=12
  o:="*"
  switch o {
  case "+": fmt.Print(a+b)
  case "-": fmt.Print(a-b)
  case "*": fmt.Print(a*b)
  case "/": fmt.Print(a/b)
  default: fmt.Print("wrong choice")

}
}
