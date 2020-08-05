package main

import (
	"io"
	"net/http"
)
func main(){
	http.HandleFunc("/",elon)
	http.ListenAndServe(":8080", nil)
}

func elon(w http.ResponseWriter,r *http.Request){
	//we are setting the header type to text html
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	//with io.writestring we are just writing the string to our server just like passing data.
	io.WriteString(w,`
<img src="https://techcrunch.com/wp-content/uploads/2020/02/GettyImages-1175368064.jpg?w=1390&crop=1">

	`)
}