package main
import("log";"text/template";"os")
func main(){
  tpl,err := template.ParseFiles("index.html")
  if err != nil {
    log.Fatalln(err)
  }
  er := tpl.Execute(os.Stdout,nil)
  if er != nil {
    log.Fatalln(err)
  }
}
