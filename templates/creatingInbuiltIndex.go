package main
import (
  "fmt"
  "strings"
  "io"
  "log"
  "os"
)
func main() {
  name := os.Args[1]  //passing data using command line
  s := fmt.Sprint(`
    <!DOCTYPE html>
  <html lang="en" dir="ltr">
    <head>
      <meta charset="utf-8">
      <title>basic</title>
    </head>
    <body>
          <h1>`+name+`</h1>
    </body>
  </html>
  `)
  newfile,err := os.Create("creatingInBuiltFile.html")
  if err != nil {
    log.Fatalln(err)
  }
  defer newfile.Close()
  io.Copy(newfile, strings.NewReader(s))
}
