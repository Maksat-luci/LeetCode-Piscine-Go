package student 

type Node struct {
	Data TreeNode
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func Print(l *List, data *TreeNode) {
	n := &Node{Data: *data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	List := &List{}
	Print(List, root)

	for l := List.Head; l != nil; l = l.Next {
		f(l.Data.Data)
		if l.Data.Left != nil {
			Print(List, l.Data.Left)
		}
		if l.Data.Right != nil {
			Print(List, l.Data.Right)
		}
	}
}
