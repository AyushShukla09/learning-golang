package main

import("log";"os";"text/template")
        //package.type
var tpl *template.Template
func init(){  //Must takes in pointer to a template and an error which is
  tpl = template.Must(template.ParseGlob("templates/*")) //output  of Glob.
}     //Must return pointer to a template..

func main(){
  err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
