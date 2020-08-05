package main

import(
  "log"
  "text/template"
  "os"
)

var tpl *template.Template

type s struct{
  Name string
  Age int
  DreamCar string
}

func init(){
  tpl = template.Must(template.ParseFiles("passingBasicData.gohtml"))
}

func main(){
  b := s{
    "Ayush Shukla",
    23,
    "Lamborghini Centanario",
  }
  err := tpl.Execute(os.Stdout,b)
  if err != nil {
    log.Fatalln(err)
  }
}
