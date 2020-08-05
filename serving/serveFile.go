package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/",em)
	http.HandleFunc("/text.txt",text)
	http.ListenAndServe(":8080",nil)
}

func em(w http.ResponseWriter,r *http.Request){
	//setting the header
	w.Header().Set("Content-Type","text/html")
	io.WriteString(w,`text.txt`)
}

func text(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"text.txt")
}




