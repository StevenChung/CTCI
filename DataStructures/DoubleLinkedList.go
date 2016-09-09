package main

import (
  "fmt"
  "errors"
)

type Value struct {
  Name string
  Numpy int
}

type Node struct {
  Value
  Next, Previous *Node // pointer to the address of a Node
}

type DoublyLinkedList struct {
  Head, Tail *Node
  Length int
}

func (l *DoublyLinkedList) First() *Node {
  return l.Head
}

func (l *DoublyLinkedList) Last() *Node {
  return l.Tail
}

func (l *DoublyLinkedList) AddToHead(v Value) *DoublyLinkedList {
  // we don't use a pointer to Value because we do not need to modify
  // the value once passed (or that's how I interpret this approach)

  newNode := &Node{Value: v}
  // Pointer to a Node that has the Value v
  // has type *Node
  // & prefix returns a pointer to the struct value

  // new allocates memory with all nil values
  if l.Head == nil {
    l.Head = newNode
    l.Tail = newNode
    // if the Head is still nil, set head to newNode
    // and subsequently increment Length upwards
  } else {
    oldHead := l.Head
    oldHead.Previous, newNode.Next = newNode, oldHead
    l.Head = newNode
  }
  l.Length++
  return l
}

func (l *DoublyLinkedList) AddToTail(v Value) *DoublyLinkedList {
  newNode := &Node{Value: v}

  if l.Tail == nil {
    l.Tail = newNode
    l.Head = newNode
  } else {
    oldTail := l.Tail
    oldTail.Next, newNode.Previous = newNode, oldTail
    l.Tail = newNode
  }
  l.Length++
  return l
}

func (l *DoublyLinkedList) RemoveFromTail() (*DoublyLinkedList, error) {
  var err error

  if l.Tail == nil {
    err = errors.New("Tail is not set")
  } else {
    newTail := l.Tail.Previous
    newTail.Next = nil
    l.Tail = newTail
    l.Length--
  }
  return l, err
}

func (l *DoublyLinkedList) RemoveFromHead() (*DoublyLinkedList, error) {
  var err error

  if l.Head == nil {
    err = errors.New("Head is not set")
  } else {
    newHead := l.Head.Next
    newHead.Previous = nil
    l.Head = newHead
    l.Length--
  }
  return l, err
}

func (l *DoublyLinkedList) DeleteFromIndex(n int) (*DoublyLinkedList, error) {
  var err error
  if n > (l.Length - 1) || n < 0 {
    err = errors.New("Invalid Index")
  } else if n == 0 || n == (l.Length - 1) {
    err = errors.New("Use the appropriate removeFromHead or RemoveFromTail")
  } else if n < (l.Length - 1 - n) {
    // start from head
    nodeToDelete := l.Head
    for i := 0; i < n; i++ {
      nodeToDelete = nodeToDelete.Next
    }
    nodeToDelete.Previous.Next, nodeToDelete.Next.Previous = nodeToDelete.Next, nodeToDelete.Previous
    l.Length--
  } else {
    // start from tail
    nodeToDelete := l.Tail
    for i := (l.Length - 1 - n); i > 0; i-- {
      nodeToDelete = nodeToDelete.Previous
    }
    nodeToDelete.Previous.Next, nodeToDelete.Next.Previous = nodeToDelete.Next, nodeToDelete.Previous
    l.Length--
  }
  return l, err
}

func (l *DoublyLinkedList) InsertAtIndex(n int, v Value) (*DoublyLinkedList, error) {
  var err error
  newNode := &Node{Value: v}

  if n > (l.Length - 1) || n < 0 {
    err = errors.New("Invalid Index")
    fmt.Println(err)
  } else if n == 0 || n == (l.Length - 1) {
    err = errors.New("Use the appropriate addToHead or addToTail")
    fmt.Println(err)
  } else if n > (l.Length - 1 - n) {
    // start from tail
    nodeToBeNext := l.Tail
    for i := (l.Length - 1 - n); i > 0; i-- {
      nodeToBeNext = nodeToBeNext.Previous
    }
    nodeToBeBefore := nodeToBeNext.Previous
    newNode.Next, newNode.Previous = nodeToBeNext, nodeToBeBefore
    nodeToBeBefore.Next, nodeToBeNext.Previous = newNode, newNode
    l.Length++
  } else {
    // start from head
    nodeBefore := l.Head
    for i := 0; i < n - 1; i++ {
      nodeBefore = nodeBefore.Next
    }
    nodeAfter := nodeBefore.Next
    newNode.Previous, newNode.Next = nodeBefore, nodeAfter
    nodeBefore.Next, nodeAfter.Previous = newNode, newNode
    l.Length++
  }
  return l, err
}

func (l *DoublyLinkedList) RetrieveAtIndex(n int) (*Node, error) {
  var err error
  var nodeToRetrieve *Node

  if n < 0 || n > (l.Length - 1) {
    err = errors.New("Invalid Index")
  } else if n < (l.Length - n) {
    nodeToRetrieve = l.Head
    for i := 0; i < n; i++ {
      nodeToRetrieve = nodeToRetrieve.Next
    }
  } else {
    nodeToRetrieve = l.Tail
    for i := l.Length - 1; i > n; i-- {
      nodeToRetrieve = nodeToRetrieve.Previous
    }
  }

  return nodeToRetrieve, err
}

func (l *DoublyLinkedList) Contains(v Value) bool {
  currentNode := l.Head
  for currentNode != nil {
    if currentNode.Value.Name == v.Name && currentNode.Value.Numpy == v.Numpy {
      return true
    }
    currentNode = currentNode.Next
    // for while loops, remember to iterate upwards
    // also, remember about the last item!
  }
  return false
}

func (l *DoublyLinkedList) CircularListCheck() bool {
  slowRunner := l.Head
  fastRunner := l.Head.Next

  for slowRunner != nil && fastRunner != nil {
    if (slowRunner.Value.Name == fastRunner.Value.Name) && (slowRunner.Value.Numpy == fastRunner.Value.Numpy) {
      return true
    }
    slowRunner = slowRunner.Next
    fastRunner = fastRunner.Next.Next
  }

  return false
}

func main() {
  newList := new(DoublyLinkedList)
  newList.AddToHead(Value{"steven", 2})
  newList.AddToHead(Value{"jaina", 1})
  newList.AddToTail(Value{"thrall", 3})
  newList.AddToTail(Value{"rexxar", 5})
  newList.AddToTail(Value{"uther", 6})

  newList.InsertAtIndex(2, Value{"gondar", 7})
  newList.InsertAtIndex(4, Value{"voljin", 17})
  newList.DeleteFromIndex(2);
  newList.DeleteFromIndex(4);

  fmt.Println(newList.CircularListCheck())
  newList.Last().Next = newList.First()
  fmt.Println(newList.CircularListCheck())

  for i := 0; i < newList.Length; i++ {
    fmt.Println(newList.RetrieveAtIndex(i))
  }
}
// Check to see if a LinkedList is a palindrome
// Reverse a LinkedList iteratively and recursively (presumably single?)
