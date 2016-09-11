/*
1.5) One Away: There are three types of edits to a given string: insert a character,
 remove a character, or replace a character.
Given two strings, write a function to check if they are one edit away (or zero).

I: Two strings
O: Boolean
C: Time/Space => None.  However, it doesn't hurt to offer the optimal solution.
E:

Are they the same string?
If so, return true.

Space?
Case sensitivity?
Empty string and a letter?
Mismatch in middle of a string vs mismatch at the end
*/

package main

import (
  "fmt"
  "math"
)

func OneAway (a, b string) bool {
  if a == b {
    return true
  } else if math.Abs(float64(len(a) - len(b))) > 1.0 {
    return false
  } else if len(a) == len(b) {
    oneFlaw := false
    runeB := []rune(b)
    for index, val := range(a) {
      if (runeB[index] != val) && (oneFlaw == false) {
        oneFlaw = true
      } else if (runeB[index] != val) && (oneFlaw) {
        return false
      }
    }
    return true
  } else {
    oneFlaw := false
    longerString := []rune{}
    shorterString := []rune{}

    if len(a) > len(b) {
      longerString = []rune(a)
      shorterString = []rune(b)
    } else {
      longerString = []rune(b)
      shorterString = []rune(a)
    }

    for i, j := 0, 0; i < len(longerString); i, j = i + 1, j + 1 {
      // i is pointer for longer string
      // j is pointer for shorter sting

      if j > (len(shorterString) - 1) {
        return true
      }

      if (longerString[i] != shorterString[j]) && (oneFlaw == false) {
        oneFlaw = true
        j--
        fmt.Println(j)
      } else if (longerString[i] != shorterString[j]) && (oneFlaw) {
        return false
      }
      // consider insertion/delete as roughly the same, but where it occurs is why we have
      // two scenarios above
      // inserting/deleting in the middle is different from the ends
    }
    return true
  }
}
func main() {
  fmt.Println(OneAway("tacty", "ttcy"))
}

// Time: O(n) Space: O(n)
