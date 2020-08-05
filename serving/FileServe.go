package main

import (
	"io"
	"net/http"
)

func main(){
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.HandleFunc("/elon/",elonM)
	http.ListenAndServe(":8080",nil)
}

func elonM(w http.ResponseWriter,r *http.Request){
	//setting the header
	w.Header().Set("Content-Type","text/html")
	//writing on website
	io.WriteString(w,`<img scr="/elonmusk.jpg">`)
}







