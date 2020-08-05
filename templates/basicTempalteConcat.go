package main
import "fmt"
func main() {
  name := "Ayush Shukla"
  s := `<!DOCTYPE html>
  <html lang="en" dir="ltr">
    <head>
      <meta charset="utf-8">
      <title>basic</title>
    </head>
    <body>
          <h1>`+name+`</h1>
    </body>
  </html>`
  fmt.Print(s)
}
//command line ====>>> go run basicTemplateConcat.go >anyname.html
