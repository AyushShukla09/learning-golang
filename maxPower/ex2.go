package main

import (
  "fmt";"math";"strconv"
  "os"
)

func maxPow(M,N float64)int{
  var k int
  for i:=0;true;i++ {
    if math.Pow(N,float64(i))<M{
      continue
    }else{
      k=i
      break
    }
  }
  return k
}
func main(){
  var err error
  var M int
  var N int
  M,err=strconv.Atoi(os.Args[1])
  N,err=strconv.Atoi(os.Args[2])
  m:=float64(M)
  n:=float64(N)
  a:=maxPow(m,n)
  fmt.Print(a)
  if err==nil{
    fmt.Print()
  }
}
