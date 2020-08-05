package main

import (
	"io"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/",root)
	http.HandleFunc("/elonmusk.jpg",emusk)
	http.ListenAndServe(":8081",nil)
}

func root(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html")
	io.WriteString(w,`<img src="/elonmusk.jpg">`)
}

func emusk(w http.ResponseWriter,r *http.Request){
	f,err:= os.Open("elonmusk.jpg")
	if err != nil {
		io.WriteString(w,"hag diya bhosadika")
	}
	defer f.Close()
	fs,err:= f.Stat() //returns file info structure
	if err != nil {
		io.WriteString(w, "laad nahi mila string")
	}
	http.ServeContent(w,r,f.Name(),fs.ModTime(),f)
	//ServeContent takes multiple arguments one of which is
	//last modification time fs.Modtime()
}
