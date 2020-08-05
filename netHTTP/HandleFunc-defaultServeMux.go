package main

import (
	"fmt"
	"net/http"
)

func d(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w,"I am a doogo")
}

func c(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w,"I am a caato")
}

func main(){
	http.HandleFunc("/dog/",d)
	http.HandleFunc("/cat/",c)
	http.ListenAndServe(":8080",nil)
}











