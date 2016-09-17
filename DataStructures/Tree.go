package main

import (
  "fmt"
)

type Node struct {
  Val int
  Children []*Node
}

type Tree struct {
  Root *Node
}

func main() {
  newTree := &Tree{Root: &Node{Val: 5}}
  newTree.Root.Children = append(newTree.Root.Children, &Node{Val: 17})
  fmt.Println(newTree.Root)
}


/*
The most basic implementation
There is one root node
Every node has zero or more children
Binary Tree is specific in that each node can have zero to two children
"Leaf" node is one that does not have children

Binary Search tree is a specific binary tree in which every node must have the property
All left descendants <= n < all right descendants (where n is a node)

     8
  4     7
2   6 7   20

Binary Search Tree

      8
    4   10
  2  12    20
  NOT a binary search tree because node 12 is on the left descendant of 8
  Must be tree for all descendants


Balanced vs Unbalanced
Slightly different definitions here, but the ideal is that it ensures O(log n) for
insert and find, but it is not the case that it is always perfectly balanced

Complete Binary Tree
A tree in which every level is filled, except possibly the last level
Must be filled left to right
Thus, the last level must be filled form left to right, (no gaps)

Full Binary Tree
Every node has zero or two children

Perfect Binary Tree
Complete and Full together!

Binary Tree Traversal!
*/
