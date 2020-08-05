package main

import(
  "log"
  "text/template"
  "os"
)

var tpl *template.Template

func init(){
  tpl = template.Must(template.ParseFiles("passingBasicData.gohtml"))
}

func main(){
  s := []string{"canada","france","venice","italy","spain"}
  err := tpl.Execute(os.Stdout,s)
  if err != nil {
    log.Fatalln(err)
  }
}
