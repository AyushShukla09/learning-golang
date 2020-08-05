package main
import("log";"text/template";"os")
func main(){
  tpl,err := template.ParseFiles("passingBasicDataInATemplate.html")
  if err != nil {
    log.Fatalln(err)
  }
  er := tpl.Execute(os.Stdout,"Ayush Shukla")
  if er != nil {
    log.Fatalln(err)
  }
}
