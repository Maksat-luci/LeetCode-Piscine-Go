package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type LList struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *LList, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func ListRemoveIf(l *LList, data_ref interface{}) {
	tmp := &LList{}
	for i := l.Head; i != nil; i = i.Next {
		if i.Data != data_ref {
			ListPushBack(tmp, i.Data)
		}
	}
	l.Head = tmp.Head
}
 