package linked_list

type (
	List interface {
		Delete(node *Node)
		Add(node *Node)
		GetLRU() *Node
	}

	list struct {
		head *Node
		tail *Node
	}

	Node struct {
		Key  string
		Val  string
		Prev *Node
		Next *Node
	}
)

func New() List {
	head := &Node{
		Val: "",
	}
	head.Next = &Node{
		Val:  "",
		Prev: head,
	}

	return &list{
		head: head,
		tail: head.Next,
	}
}

func (l *list) Delete(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}

func (l *list) Add(node *Node) {
	tmp := l.head.Next
	l.head.Next = node

	node.Prev = l.head
	node.Next = tmp

	tmp.Prev = node
}

func (l *list) GetLRU() *Node {
	return l.tail.Prev
}
