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

// 在index处新增，原Index位置的元素后移
func (list *List) Insert(data Object, index int) {
	list.size++
	if index < 0 {
		list.AddFirst(data)
	} else if index > list.size {
		list.Append(data)
	} else {
		pre := list.headNode
		node := &Node{
			Data: data,
		}
		if pre == nil {
			list.headNode = node
			return
		}
		if index == 0 {
			node.Next, list.headNode = pre, node
			return
		}
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		// 此时pre指向index-1
		node.Next, pre.Next = pre.Next, node
	}
}

func (list *List) Remove(data Object) {
	list.size--
	pre := list.headNode
	if pre.Data == data {
		list.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			}
		}
	}
}

func (list *List) RemoveAtIndex(index int) {
	list.size--
	pre := list.headNode
	if index <= 0 {
		list.headNode = pre.Next
	} else if index > list.size {
		return
	} else {
		count := 0
		for count != (index-1) && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

func (list *List) get(index int) *Node {
	if index < 0 || index > list.size {
		panic("over limit")
	}
	count := 0
	cur := list.Head()
	for count < index {
		cur = cur.Next
		count++
	}
	return cur
}

func (list *List) Head() *Node {
	return list.headNode
}

////////////// Node /////////////////
