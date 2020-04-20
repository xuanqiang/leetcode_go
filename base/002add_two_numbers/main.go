/*
 * 给出两个 非空 的链表用来表示两个非负的整数。
 * 其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储一位数字。
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 */
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	//定义头节点
	result := list
	temp := 0
	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{temp % 10, nil}
		temp = temp / 10
		list = list.Next
	}
	return result.Next
}

//使用头插法建立
//要想通过函数改变外面的变量的值 必须传进去这个链表地址的地址
func insertLinkList(list **ListNode, val int) {
	node := &ListNode{0, nil} //创建一个新节点
	node.Val = val
	node.Next = *list
	*list = node
}

//打印链表
func printLinkList(list *ListNode) {
	for list != nil && list.Next != nil {
		fmt.Print(list.Val)
		list = list.Next
	}
	fmt.Println()
}

func main() {
	//建立两个逆序链表
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}
	arr1 := [3]int{2, 4, 3}
	arr2 := [3]int{5, 6, 4}
	for _, v := range arr1 {
		insertLinkList(&l1, v)
	}

	for _, v := range arr2 {
		insertLinkList(&l2, v)
	}
	result := addTwoNumbers(l1, l2)
	printLinkList(result)
}
