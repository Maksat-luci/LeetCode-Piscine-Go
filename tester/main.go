package main

import (
	"fmt"
	"strings"
	"student"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type LList struct {
	Head *NodeL
	Tail *NodeL
}

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func Printlist(l *student.LList) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
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

func main() {
	fmt.Println(student.Atoi("9223372036854775807"))
	fmt.Println("ATOI |||")
	fmt.Println()
	fmt.Println(strings.Split("HAHA", "HA"))
	fmt.Println()
	fmt.Println(student.Split("HAHA", "HA"))
	fmt.Println("SPLIT |||")
	fmt.Println()
	fmt.Println()
	fmt.Println(student.ActiveBits(7))
	fmt.Println("ACTIVEBITS |||")
	fmt.Println()
	result := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	f := func(a, b string) int {
		return strings.Compare(b, a)
	}
	student.AdvancedSortWordArr(result, f)
	fmt.Println(result)
	fmt.Println("ADVANCEDSORTWORDARR |||")
	fmt.Println()
	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println()
	fmt.Println("ATOIBASE |||")
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeApplyByLevel(root, fmt.Println)
	fmt.Println("btreeapplybelvl |||")
	fmt.Println()

	root2 := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root2, "1")
	student.BTreeInsertData(root2, "7")
	student.BTreeInsertData(root2, "5")
	node := student.BTreeSearchItem(root2, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root2, fmt.Println)
	root2 = student.BTreeDeleteNode(root2, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root2, fmt.Println)
	fmt.Println("btreedeletenode |||")
	fmt.Println()
	root3 := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root3, "1")
	student.BTreeInsertData(root3, "7")
	student.BTreeInsertData(root3, "5")
	node2 := student.BTreeSearchItem(root3, "1")
	replacement := &student.TreeNode{Data: "3"}
	root3 = student.BTreeTransplant(root3, node2, replacement)
	student.BTreeApplyInorder(root3, fmt.Println)
	fmt.Println("btreetransplant |||")
	fmt.Println()
	fmt.Println(student.ConvertBase("683241", "0123456789", "0123456789"))
	fmt.Println("CONVERTBASE |||")
	fmt.Println()
	link666 := &student.LList{}
	link777 := &student.LList{}

	fmt.Println("----normal state----")
	student.ListPushBack(link777, 1)
	Printlist(link777)
	student.ListRemoveIf(link777, 1)
	fmt.Println("------answer-----")
	Printlist(link777)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link666, 1)
	student.ListPushBack(link666, "Hello")
	student.ListPushBack(link666, 1)
	student.ListPushBack(link666, "There")
	student.ListPushBack(link666, 1)
	student.ListPushBack(link666, 1)
	student.ListPushBack(link666, "How")
	student.ListPushBack(link666, 1)
	student.ListPushBack(link666, "are")
	student.ListPushBack(link666, "you")
	student.ListPushBack(link666, 1)
	Printlist(link666)

	student.ListRemoveIf(link666, 1)
	fmt.Println("------answer-----")
	Printlist(link666)
	fmt.Println("LISTREMOVEIF |||")
	fmt.Println()
	student.PrintCombN(1)
	student.PrintCombN(3)
	fmt.Println("PRINTCOMBN |||")
	fmt.Println()
	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	fmt.Println()
	fmt.Println("PRINTNBRBASE |||")
	fmt.Println()
	fmt.Println(student.RecursivePower(-6, 5))
	fmt.Println("RECURSIVEPOWER |||")
	fmt.Println()
	var link *student.NodeI
	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)
	PrintList(link)
	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)
	fmt.Println("SORTLISTINSERT |||")
	fmt.Println()
	var link2 *student.NodeI
	var link3 *student.NodeI
	link2 = listPushBack(link2, 3)
	link2 = listPushBack(link2, 5)
	link2 = listPushBack(link2, 7)
	link3 = listPushBack(link3, -2)
	link3 = listPushBack(link3, 9)
	PrintList(student.SortedListMerge(link3, link2))
	fmt.Println("SORTLISTMERGE |||")
	fmt.Println()
	fmt.Println(strings.Split("HAHA", "HA"))
	fmt.Println()
	fmt.Println(student.Split("HAHA", "HA"))
	fmt.Println("SPLIT |||")
	fmt.Println()
	fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	fmt.Println("SPLITWHITESPACES |||")
	fmt.Println()

}
