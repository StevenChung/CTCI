/*
The most basic implementation
There is one root node
Every node has zero or more children
Binary Tree is specific in that each node can have zero to two children
"Leaf" node is one that does not have children

Tree
Binary Tree (Zero to Two Nodes)
- Balanced vs Unbalanced
a) O(log n) for insert/find
- Complete
Every level is filled except possibly the last
More balanced!
- Full
Every node has zero or two nodes
- PERFECT
Combination of the above
-Binary Search Tree
All left descendants <= n < all right descendants (where n is a node)
-Min/Max Heap
Where each node is less than or greater than its immediate children; thus, the root is the largest/smallest value
a)
IMPORTANT: BINARY SEARCH TREE
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

MIN HEAP
COMPLETE binary tree (not necessarily full) each node is less than its children
Each node is less than its immediate children
Thus, the root is the minimum element in the tree
INSERT
and EXTRACT_MIN (or MAX)
are the two functions



Full Binary Tree
Every node has zero or two children

Perfect Binary Tree
Complete and Full together!

Binary Tree Traversal!

      2
    1   4
  0   1.1 3  5



- In-order Traversal

func (t *Tree) inOrderTraversal(n *Node) {
  if n != nil {
    t.inOrderTraversal(n.Left)
    fmt.Println(n)
    t.inOrderTraversal(n.Right)
  }
}

0 1 1.1 2 3 4 5
- Pre-Order Traversal

func (t *Tree) preOrderTraversal(n *Node) {
  if n != nil {
    fmt.Println(n)
    t.preOrderTraversal(n.Left)
    t.preOrderTraversal(n.Right)
  }
}

2 1 0 1.1 4 3 5

- Post-Order Traversal

func (t *Tree) postOrderTraversal(n *Node) {
  if n != nil {
    t.postOrderTraversal(n.Left)
    t.postOrderTraversal(n.Right)
    fmt.Println(n)
  }
}

0 1.1 1 3 5 4 2

*/

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
