/*
Objec­tive: Given a Binary tree cre­ate Linked Lists of all the nodes at each depth , say if the tree has height k then cre­ate k linked lists.

List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth
(e.g., if you have a tree with depth x, you'll have x linked lists).

I: Binary Tree (maybe not BST)
O: Linked Lists (n) where n is also the depth
C: None
E:
Unbalanced Tree?

Just one branch?
- a bunch of one node linkedLists?

Root is just one node?
*/

package main

import (
  "fmt"
)

type LinkedListNode struct {
  Val int
  Next *LinkedListNode
}

type BinaryTreeNode struct {
  Val int
  Left *BinaryTreeNode
  Right *BinaryTreeNode
}

// func (bst *BinaryTreeNode) LinkedListDepth() []*LinkedListNode {
//
// }
func inOrderTraversal(btn *BinaryTreeNode) {
  // does not need to be a pointer method since we're just printing nodes
  if btn != nil {
    inOrderTraversal(btn.Left)
    fmt.Println(btn.Val)
    inOrderTraversal(btn.Right)
  }
}

func main() {
  bst := &BinaryTreeNode{Val: 1}
  // create a pointer to BinaryTreeNode that has value of 1
  // new is equivalent to this
  // var bst *BinaryTreeNode is different in that it offers a nil value
  // pointer and new saves space in memory with nil values (different!)
  bst.Left = &BinaryTreeNode{Val: 2}
  bst.Right = &BinaryTreeNode{Val: 3}
  bst.Left.Left = &BinaryTreeNode{Val: 4}
  bst.Left.Right = &BinaryTreeNode{Val: 5}
  bst.Right.Left = &BinaryTreeNode{Val: 6}
  bst.Right.Right = &BinaryTreeNode{Val: 7}
  inOrderTraversal(bst)
}
