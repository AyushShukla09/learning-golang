package main
import ("log";"text/template";"strings";"os")

var tpl *template.Template

var fn = template.FuncMap{//FuncMap underlying type is string and returns empy interface
  "uc" : strings.ToUpper, //fn should be declared before it is used during init()
  "rev" : reverse,
}

func init(){
  tpl = template.Must(template.New("").Funcs(fn).ParseFiles("passingBasicData.gohtml"))
}//Funcs take in parameter of type FuncMap
//Funcs need to be there before the tempates are passed thats why Funcmap is above
//and before parseFiles.


 func reverse(s string)string{
   t := ""
   for _,v:= range s{
     t=string(v)+t
   }
   return t
 }

func main(){
  s := []string{"ayush","shukla"}
  err := tpl.ExecuteTemplate(os.Stdout,"passingBasicData.gohtml",s)
  //funcmap uses only ExecuteTemplate to execute not just Execute
  if err != nil{
    log.Fatalln(err)
  }
}
