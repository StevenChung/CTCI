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
*/

package main

import (
  "fmt"
  "math"
)

func OneAway (a, b string) bool {
  if a == b {
    return true
    // if they're the same string, it's true!
  } else if math.Abs(float64(len(a) - len(b))) > 1 {
    // if the difference in length is greater than one, removing, adding or
    // replacing one character will not suffice, so it must be false
    return false
  } else {
    runeHashA := map[rune]int{}
    runeHashB := map[rune]int{}
    // to initiate a map, need to either use make OR
    // type inference assignment with curly braces

    for _, runevalue := range a {
      // range on a string offers the index and the RUNE value
	     if runeHashA[runevalue] == 0 {
         runeHashA[runevalue] = 1
       } else {
         runeHashA[runevalue] += 1
       }
    }
    for _, runevalueb := range b {
      // range on a string offers the index and the RUNE value
       if runeHashB[runevalueb] == 0 {
         runeHashB[runevalueb] = 1
       } else {
         runeHashB[runevalueb] += 1
       }
    }

    flawcounter := 0
    for k, _ := range runeHashA {
      if runeHashB[k] == 0 || runeHashA[k] != runeHashB[k] {
        flawcounter++
      }
      if flawcounter > 1 {
        return false
      }
    }
    return true

  }
}
func main() {
  fmt.Println(OneAway("dag", "dog"))
}

// Time: O(n) Space: O(n)
