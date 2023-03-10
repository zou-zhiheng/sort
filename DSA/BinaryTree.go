package DSA

import "fmt"

type binaryTree struct {
	val   int
	left  *binaryTree
	right *binaryTree
}


//二叉树创建
func createBinaryTree(tree **binaryTree) {
	n := 0
	_, _ = fmt.Scanf("%d", &n)
	if n == 0 {
		*tree = nil
	}else{
		*tree =new(binaryTree)
		(*tree).val=n
		createBinaryTree(&(*tree).left)
		createBinaryTree(&(*tree).right)
	}

}


//前序遍历，先输出父节点，再遍历左子树和右子树
//中序遍历，先遍历左子树，再输出父节点，再遍历右子树
//后序遍历，先遍历左子树，在遍历右子树，最后输出父节点


//前序遍历
func preOrder(node *binaryTree) {
	if node != nil {
		fmt.Println(node.val)
		preOrder(node.left)
		preOrder(node.right)
	}
}

//中序遍历
func infixOrder(node *binaryTree) {
	if node != nil {
		preOrder(node.left)
		fmt.Println(node.val)
		preOrder(node.right)
	}
}

func postOrder(node *binaryTree) {
	if node != nil {
		preOrder(node.left)
		preOrder(node.right)
		fmt.Println(node.val)
	}
}

func BinaryTreeDemo(){
	tree:=binaryTree{val: 1}
	createBinaryTree(&tree.left)
	createBinaryTree(&tree.right)
	fmt.Println(tree)
	//preOrder(&tree)
	infixOrder(&tree)
}
