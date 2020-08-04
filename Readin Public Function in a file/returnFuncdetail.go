package main

var kk ="func"
func main(){
	aa:="func hello( s str) int {"
	Disp(aa)
}
func Disp(st string) {
	es:=""
	ss:=""
	for _,v := range st{
		if es==kk && string(v) != "{" {
			ss+=string(v)		// hello( s str) int
			continue
		}else if string(v) == "{"{
			es+=string(v)				//func{
			print(ss[1:])//_hello( s str) int_
			return
		}else{
			es+=string(v)//func
		}
	}
}



