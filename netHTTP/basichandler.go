package main

import (
	"net/http"
)

type hotdo int //creates a type hotdog whose underlying type is int

func (h hotdo) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
/*
func (m d)ServeHTTP(w http.ResponseWriter,r *http.Request){ //attach this handler ServeHTTP to type hotdog
	fmt.Fprintln(w,"mahnat kro bhai") //writing a string to a output
}

 */

func main(){
	var d hotdo //creating a variable d of type hotdog whose underlying type is int and a type handler attached to it
	http.ListenAndServe(":8080",d) //whenever a request comes at port 8080 handle the request with handler d.
}

