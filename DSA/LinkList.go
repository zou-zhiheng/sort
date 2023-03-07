package DSA

import (
	"errors"
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode
}

//添加节点，尾插法
func (r *HeroNode) addLinkList(node HeroNode) {
	//找到尾节点
	l := r
	for ; l.next != nil; l = l.next {
	}
	l.next = &node

}

//前插法
func addLinkList(node HeroNode, L *HeroNode) {
	if L == nil {
		L = &node
	} else {
		pre := &node
		pre.next = L
	}
}

//链表是否为空
func (r *HeroNode) isEmpty() bool {

	if r == nil {
		return true
	}

	return false
}

//节点删除
func deleteLinkList(no int, r *HeroNode) error {

	pre, l := r, r
	for ; l != nil; l = l.next {
		//fmt.Println(l.no)
		if l.no == no {
			if pre == l { //头结点
				r = r.next
			} else if l.next == nil { //尾节点
				pre.next = nil
			} else {
				pre.next = l.next
			}

			return nil
		}
		pre = l
	}

	return errors.New("没有找到该节点")
}

func reverse(L *HeroNode) *HeroNode {
	var pre, curr, next *HeroNode //新头结点，要插入的新节点，新链表头节点后的节点
	for ; L != nil; {
		curr = L   //要新插入的节点
		L = L.next //循环遍历
		next = pre //保存新链表头节点后的节点
		curr.next = nil
		pre = curr      //新头节点
		pre.next = next //连接之前的节点
	}

	return pre

}

func HeroNodeDemo() {
	test := &HeroNode{}
	test.addLinkList(HeroNode{no: 1})
	test.addLinkList(HeroNode{no: 2})
	test.addLinkList(HeroNode{no: 3})

	//p := test
	//for ; test != nil; test = test.next {
	//	fmt.Println(test.no)
	//}
	//test = reverse(p)
	//fmt.Println("show------")
	p:=test
	//for ; test != nil; test = test.next {
	//	fmt.Println(test.no)
	//}
	showReverse(p)

}

func showReverse(L *HeroNode) {

	if L.next == nil {
	} else {
		showReverse(L.next)
	}
	fmt.Println(L.no)
}
