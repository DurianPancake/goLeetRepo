package utils

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
	size     int
}

func (list *List) IsEmpty() bool {
	return list.headNode == nil
}

func (list *List) Length() int {
	return list.size
}

//
func (list *List) AddFirst(data Object) *Node {
	node := &Node{
		Data: data,
	}
	node.Next, list.headNode = list.headNode, node
	list.size++
	return node
}

func (list *List) Append(data Object) {
	node := &Node{
		Data: data,
	}
	list.size++
	if list.IsEmpty() {
		list.headNode = node
	} else {
		cur := list.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (list *List) Insert(data Object, index int) {
	list.size++
	if index < 0 {
		list.AddFirst(data)
	} else if index > list.size {
		list.Append(data)
	} else {
		pre := list.headNode
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		// 此时pre指向index-1
		node := &Node{
			Data: data,
		}
		node.Next, pre.Next = pre.Next, node
	}
}

func (list *List) Head() *Node {
	return list.headNode
}

func (node *Node) HasNext() bool {
	return node.Next != nil
}

func (node *Node) getNext() *Node {
	return node.Next
}
