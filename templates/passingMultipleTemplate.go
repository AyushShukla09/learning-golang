package main
import ("log";"os";"text/template")
func main() {
  tpl,err := template.ParseFiles("abc.html")
  if err != nil {
    log.Fatalln(err)
    }
    err = tpl.Execute(os.Stdout,nil) //will execute the last passes template.
    if err != nil {
      log.Fatalln(err)
      }
  tpl,err = tpl.ParseFiles("index.html","creatingInBuiltFile.html")
  if err != nil {
    log.Fatalln(err)
    }
  err = tpl.ExecuteTemplate(os.Stdout,"creatingInBuiltFile.html", nil)
  if err != nil {
    log.Fatalln(err)
    }
  tpl,err = tpl.ParseFiles("basicHTML.html") //the last passes template execute.
  if err != nil {
    log.Fatalln(err)
    }

//to execute a specific template say index.html
  err = tpl.ExecuteTemplate(os.Stdout,"basicHTML.html", nil)
  if err != nil {
    log.Fatalln(err)
    }

  err = tpl.Execute(os.Stdout,nil) //will execute the last passes template.
  if err != nil {
    log.Fatalln(err)
    }



}
