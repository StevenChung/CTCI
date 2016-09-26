/*
For the purposes of this question, a balanced tree is defined to be a tree
Implement a function to check if a binary tree is balanced.
such that the heights of the two subtrees of any node
never differ by more than one

https://leetcode.com/problems/balanced-binary-tree/
*/
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func returnDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	fmt.Println(node.Value)
	left := returnDepth(node.Left)
	if left != -1 {
		right := returnDepth(node.Right)
		if right != -1 {
			absV := math.Abs(float64(left) - float64(right))
			if absV <= 1 {
				if left > right {
					return left + 1
				} else {
					return right + 1
				}
			} else {
				return -1
			}
		}
	}
	return -1
}

func isBalanced(rootNode *TreeNode) bool {
	return returnDepth(rootNode) != -1
}

func main() {
	balancedTree := &TreeNode{Value: 1}
	balancedTree.Left = &TreeNode{Value: 2}
	balancedTree.Right = &TreeNode{Value: 3}
	balancedTree.Left.Left = &TreeNode{Value: 4}
	balancedTree.Left.Right = &TreeNode{Value: 5}

	unbalancedTree := &TreeNode{Value: 1}
	unbalancedTree.Left = &TreeNode{Value: 2}
	unbalancedTree.Right = &TreeNode{Value: 3}
	unbalancedTree.Right.Right = &TreeNode{Value: 4}
	unbalancedTree.Right.Right.Right = &TreeNode{Value: 4}
	unbalancedTree.Right.Right.Right.Right = &TreeNode{Value: 4}

	fmt.Println(isBalanced(balancedTree))
	fmt.Println(isBalanced(unbalancedTree))
}
