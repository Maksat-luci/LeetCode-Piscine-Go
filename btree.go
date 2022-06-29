package student

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if res := BTreeSearchItem(root.Left, elem); res != nil {
		return res
	}
	return BTreeSearchItem(root.Right, elem)
}


func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data <= data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	} else {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	}
	return root
}


func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}