/*
https://leetcode.com/problems/add-two-numbers/
*/

package main

type Node struct {
  Value int
  Next *Node // pointer to the address of another Node
}

type LinkedList struct {
  Head *Node // pointer to the address of another Node
}

func (l *LinkedList) Append(n int) {
  newNode := &Node{Value: n}
  // newNode is set to type *Node
  // pointer to a Node that has value n
  if l.Head == nil {
    l.Head = newNode
  } else {
    currentNode := l.Head
    for currentNode.Next != nil {
      currentNode = currentNode.Next
    }
    currentNode.Next = newNode
  }
}

// func main() {
//   newList := &LinkedList{}
//   newList.Append(2)
//   newList.Append(3)
//   newList.Append(4)
//
//   currentNode := newList.Head
//   for currentNode != nil {
//     fmt.Println(currentNode)
//     currentNode = currentNode.Next
//   }
// }
