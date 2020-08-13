package main

type Queue []string

func (q *Queue)Enqueue(value string){
	*q=append(*q,value)
	println("value added to queue : ",value)
}
func (q *Queue)Dequeue(){
	if len(*q)==0{
		println("Dequeue is impossible as the queue is empty")
		return
	}
	println("poped item is : ",(*q)[0])
	*q = (*q)[1:]
}

func (q *Queue)IsFront(){
	if len(*q)==0{
		println("IsFront is impossible as the queue is empty")
		return
	}
	println("peeked item is : ",(*q)[0])
}

func (q *Queue)IsEmpty(){
	if len(*q)!=0{
		println("the queue is not empty")
		return
	}
	println("the queue is empty")
}

func (q *Queue)Traverse(){
	if len(*q)==0{
		println("Traversing not possible as the queue is empty")
		return
	}
	for i:=0;i<len(*q);i++{
		print((*q)[i]+" ")
	}
}

func main(){
	var queue Queue
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Dequeue()
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Dequeue()
	queue.IsFront()
	queue.IsEmpty()
	queue.Traverse()
}
