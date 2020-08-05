package main

import (
	"fmt"
	"net/http"
)

type hotdog int // creating a handler

func main(){
	var d hotdog //
	http.ListenAndServe(":8080",d)
}

func (m hotdog)ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {// fetches the path entered in url
	case "/dog":
		fmt.Fprint(w,"i am a doggo")
	case "/cat":
		fmt.Fprint(w,"i am a catto")
	}
}