package main

import(
	"io"
	"net/http"
	"os"  //to open source files and later copy it file writer object and then to server
)

func main(){
	http.HandleFunc("/",ownserver) 	//at root run this ownserver handlefunc
	http.HandleFunc("/elonmusk.jpg",elonmusk) //this will be called when io.WriteString will be looking for /elonmusk.jpg file
	http.ListenAndServe(":8080",nil)
}

func ownserver(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	//this write string is going to call the elonmusk handlefunc by looking for /elonmusk.jpg pattern
	io.WriteString(w,`
	<img src="/elonmusk.jpg"> 
	`)
	//when this image elonmusk.jpg is asked for then that elonmusk handlefunc will be called.
}

func elonmusk(w http.ResponseWriter,r *http.Request){
	fp,err:=os.Open("elonmusk.jpg") //open function returns a filepointer that can be used to read from/ write to a file
	if err != nil {
		http.Error(w,"file not found",404)
		return //make sure to return else in the first go will not be able get out of 404 error
	}
	defer fp.Close()

	io.Copy(w,fp) //copying the bytes of filepointer to the web server. i.e printing the image/data.
}