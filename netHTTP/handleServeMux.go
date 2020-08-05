package main

import (
	"fmt"
	"net/http"
)
type dog int

func (a dog) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w,"I am a doogo")
}

type cat int

func (b cat) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w,"I am a caato")
}

func main(){
	var a dog
	var b cat
	mux := http.NewServeMux()
	mux.Handle("/dog/",a)
	mux.Handle("/cat/",b)
	http.ListenAndServe(":8080",mux)
}











