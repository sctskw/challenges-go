package main

import "fmt"

type Node struct {
	position int
	value    string
	next     *Node
	list     *LinkedList
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) String() string {
	return fmt.Sprintf("[%d, %s]", n.position, n.value)
}

type LinkedList struct {
	len  int
	head *Node
	tail *Node
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Tail() *Node {

	if l.len <= 1 {
		return l.Head()
	}

	return l.tail
}

func (l *LinkedList) AddValue(v string) *LinkedList {
	return l.Add(&Node{value: v})
}

func (l *LinkedList) Add(n *Node) *LinkedList {

	n.list = l
	n.position = l.len

	if l.len == 0 {
		l.head = n
		l.tail = n
		l.len = 1
		return l
	}

	l.tail.next = n
	l.tail = n
	l.len = l.len + 1

	return l
}

func main() {

	list := new(LinkedList)

	list.AddValue("A")
	list.AddValue("b")
	list.AddValue("c")
	list.AddValue("d")
	list.AddValue("e")

	curr := list.Head()

	for curr != nil {
		fmt.Println(curr)
		curr = curr.Next()
	}

}
