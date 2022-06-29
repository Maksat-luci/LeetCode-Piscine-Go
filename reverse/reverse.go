package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Reverse(node *NodeAddL) *NodeAddL {
	for i := node; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Num < j.Num {
				i.Num, j.Num = j.Num, i.Num
			}

		}
	}

	return node
}

func pushBack(n *NodeAddL, num int) *NodeAddL {
	l := &NodeAddL{Num: num}
	l.Next = n
	n = l
	return l
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)
	result := Reverse(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
