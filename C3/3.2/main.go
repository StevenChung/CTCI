/*
Stack Min: How would you design a stack which, in addition to push and pop, has a function min
which returns the minimum element? Push, pop and min should all operate in 0(1) time.
*/

package main

import (
  "fmt"
  // "errors"
)

// func (s *Stack) Min() (int, error) {
//   // if we wanted O(1) time, my naive solution would be to track the min as items are appended
//   // thus, when we add, we check against the min
//   /*
//   Sequence
//   Stack has another property called minStack (which is itself a Stack)
//   Stack
//   7 -> 3 -> 9 -> 2
//   MinStack (of the above Stack)
//   7 -> 3 -> 2
//   Popping from main stack has a check to see if that node has the same value from min stack
//   if so, similarly pop it off
//   thus, min() returns the top value from a stack's minstack
//   */
// }


func main() {
  newStack := &Stack{}
  newStack.Add(5)
  newStack.Add(7)
  newStack.Add(-99)
  newStack.Add(-1)
  fmt.Println(newStack.Last)
  fmt.Println(newStack.LocalMin)
}
