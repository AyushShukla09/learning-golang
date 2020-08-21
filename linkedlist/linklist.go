package main

type Node struct {
	value int
	next  *Node
}

type list struct {
	len  int
	head *Node
	tail *Node
}

func (l *list) AddFront(v int) {
	node := &Node{
		v,
		nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.len++
	return
}

func (l *list) AddLast(v int) {
	node := &Node{
		v,
		nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		temp:=l.head
		prev:=temp
		for temp != nil{
			prev=temp
			temp=temp.next
		}
		prev.next=node
	}
	l.len++
	return
}

func (l *list) Insert(index int, v int) {
	node := &Node{
		v,
		nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		if index > l.len{
			l.AddLast(v)
			return
		}
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		node.next = current.next
		current.next = node
	}
}

func (l *list) DeleteFirst() {
	if l.len == 0 {
		println("link list is empty ")
		return
	}
	println("the deleted element from front is : ", l.head.value)
	l.head = l.head.next
	l.len--
}

func (l *list) Delete(index int) {
	if l.len == 0 {
		println("link list is empty ")
		return
	}
	temp := l.head
	if index == 0{
		println("value deleted at given index ", index, " is = ", temp.value)
		l.head=temp.next
	}else if index > l.len{
		l.DeleteLast()
		return
	}else{
		for i := 0; i < index-1; i++ {
			temp = temp.next
		}
		println("value deleted at given index ", index, " is = ", temp.next.value)
		temp.next=temp.next.next
	}
	l.len--
}

func (l *list) DeleteLast() {
	if l.len == 0 {
		println("link list is empty ")
		return
	}
	prev := l.head
	temp := l.head
	for temp.next != nil {
		prev = temp
		temp = temp.next
	}
	prev.next = nil
	println("the deleted element from back is : ", temp.value)
	l.len--
}

func (l *list) Traverse() {
	if l.len == 0 {
		println("link list is empty ")
		return
	}
	temp := l.head
	for temp != nil {
		print(temp.value, " ")
		temp = temp.next
	}
	println()
}

func main() {
	l := &list{}
	l.AddFront(5)
	l.AddLast(9)
	l.Insert(1, 51)
	l.Insert(20, 50)
	l.Insert(22, 58)
	l.Insert(27, 25)
	l.Traverse()
	l.DeleteFirst()
	l.DeleteLast()
	l.Delete(0)
	l.Traverse()
}
