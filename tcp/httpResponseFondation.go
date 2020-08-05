package main

import(
  "bufio"
  "fmt"
  "log"
  "net"
  "strings"
)

func main(){
  li,err := net.Listen("tcp",":8080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()

for {
  conn,err:=li.Accept()
  if err != nil {
    log.Fatalln(err)
    continue
  }
  go handle(conn)
 }
}

func handle(conn net.Conn){
  defer conn.Close()
//request
  request(conn)
//listen
  respond(conn)
}

func request(conn net.Conn){
  i:=0// counter made for signelling in the program
  scanner:=bufio.NewScanner(conn)
  for scanner.Scan(){
    ln := scanner.Text()
    fmt.Println(ln)
    if i==0 {
      m:=strings.Fields(ln)[0]//gives a slice of strings and return method GET which a line 0 of request line
    //  url:=strings.Fields(ln)[1]//gives the url 
      fmt.Println("MEATHOD",m)
    }
    if ln==""{//break out loop when there is nothing in the line
      //headers are done
      break
    }
    i++
  }
}

func respond(conn net.Conn){
//rudimentary way of passing html
  body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
//                                   \r is carriage return
  	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")//this is the start line which is the status line
  	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  	fmt.Fprint(conn, "Content-Type: text/html\r\n")// tells which type of is to receive
  	fmt.Fprint(conn, "\r\n") //blank line
  	fmt.Fprint(conn, body)
}
