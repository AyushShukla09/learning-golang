package main

import (
	"net/http"
	"text/template"
)

var tpo *template.Template
type dataa struct {
	AccType       string
	BranchCode    string
	Balance       string
	ContactNumber string
}
var infoo map[string]dataa //declaring a composite data structure
var flag int               //a flag varible to check if user has entered valid Account ID
var AcID string            //a variable to store user entered account ID
func init() {
	tpo,_ = template.ParseGlob("template/*") //passing the entire folder containing templates
	//making a composite data structure with key as string and a value as struct
	infoo = map[string]dataa{
		"123": {"saving", "Bengaluru", "10000", "99999999"},
		"456": {"current", "Mumbai", "20000", "888888888"},
		"789": {"investing", "Chennai", "30000", "777777777"},
	}
}
func main() {
	http.HandleFunc("/", acc)           //the root handler
	http.HandleFunc("/edit", edi)       //the edit handler
	http.HandleFunc("/accinfo", accinf) //the account info displaying handler
	http.ListenAndServe(":9000", nil)
}
//form taking account ID for Input
func acc(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	if AcID ==""{
		AcID = r.FormValue("accID")
		err = tpo.ExecuteTemplate(w, "EnterAccID.html", nil)
		if err != nil {
			panic(err)
		}
		return
	}
	return
}
//form for displaying account info
func accinf(w http.ResponseWriter, r *http.Request) {
	for k,_ := range infoo { //ranging through the saved data base to match key with entered account ID
		if k == AcID { //if user entered account ID is found in map of structs / database
			flag = 1 //update flag to tell match found
			err := tpo.ExecuteTemplate(w, "accinfo.html", infoo[AcID])
			if err != nil {
				panic(err)
			}
		}
	}
	if flag == 0 { //if no match is found
		err := tpo.ExecuteTemplate(w, "alternate.html", nil) //will display this template in case id does not match
		if err != nil {
			panic(err)
		}
	}
}
//form for editing contact
func edi(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	if r.FormValue("contactNumber") != "" { //additional check if user has entered something then only it will update
		t := infoo[AcID] //else it will display the same information
		t.ContactNumber = r.FormValue("contactNumber")
		infoo[AcID] = t
	}
	if flag == 1 { //check : if user entered a valid ACCOUNT ID
		err = tpo.ExecuteTemplate(w, "editinfo.html", nil)
		if err != nil {
			panic(err)
		}
	} else { //incase a user entered an invalid account ID
		err := tpo.ExecuteTemplate(w, "alternate.html", nil)
		if err != nil {
			panic(err)
		}
	}
}