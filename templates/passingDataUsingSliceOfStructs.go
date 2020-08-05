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
}

func init(){
  tpl = template.Must(template.ParseFiles("passingBasicData.gohtml"))
}

func main(){
  b := s{
    "Ayush Shukla",
    23,
  }
  a := s{
    "Ayush Shukla",
    23,
  }
  c := s{
    "Ayush Shukla",
    23,
  }
  d := s{
    "Ayush Shukla",
    23,
  }
  e:=[]s{a,b,c,d}
  err := tpl.Execute(os.Stdout,e)
  if err != nil {
    log.Fatalln(err)
  }
}
