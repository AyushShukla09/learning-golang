package main
import("log";"text/template";"os")
func main(){
  tpl,err := template.ParseFiles("index.html")   //we can pass multiple html files in here
  if err != nil {
    log.Fatalln(err)
  }
  nf,errr:=os.Create("abc.html")      //creating new file
  defer nf.Close()                    //closing the file
  if errr != nil {
    log.Fatalln(errr)
  }
  er := tpl.Execute(nf,nil)           //we can pass nf here as nf contains
  if er != nil {               //writer interface(os.Stdout)
    log.Fatalln(err)
  }
}
