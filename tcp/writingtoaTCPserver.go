package main

import (
  "log";"net";"io"
)

func main(){
  li,err := net.Listen("tcp",":8080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()
  for{
    conn,err := li.Accept()
    if err != nil {
      log.Fatalln(err)
    }
    defer conn.Close()
    io.WriteString(conn,"\n yes \n")
  }



}
