package main


type Stack []string

func (s *Stack)Push(str string){
	//push item on top of stack
	*s=append(*s,str)
	println("pushed an element in the stack",str)
}
//pop the top most item from stack
func (s *Stack)Pop(){
	if len(*s)==0{
		println("Pop not possible as the stack is empty")
		return
	}
	str:=(*s)[len(*s)-1] //put *s in () is important else [] will work on only s not *s
	*s=(*s)[:len(*s)-1]
	println("popped element is ",str)
}
//return topmost value without deleting it
func (s *Stack)Peek(){
	if len(*s)==0{
		println("Peek not possible as the stack is empty")
		return
	}
	str:=(*s)[len(*s)-1] //put *s in () is important else [] will work on only s not *s
	println("top most element",str)
}

//display the stack
func (s *Stack)Display(){
	if len(*s)==0{
		println("Traversing not possible as the stack is empty")
		return
	}
	for i:=0;i<len(*s);i++{
		print((*s)[i]+" ")
	}
}

func main(){
	var stack Stack
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Pop()
	stack.Pop()
	stack.Peek()
	stack.Pop()
	stack.Peek()
	stack.Pop()
	stack.Display()
}