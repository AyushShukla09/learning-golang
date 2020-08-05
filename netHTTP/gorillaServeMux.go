package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/dog/",dd)
	r.HandleFunc("/cat",cc)
	http.ListenAndServe(":9000",r)
}
func dd(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"i am a doggysur")
}
func cc(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"i am a cattysur")
}
