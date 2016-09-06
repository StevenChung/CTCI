/*
1.6) String Compression
Implement a method to perform a basic compression using the counts of
repeated characters.  For example, the string `aabcccccaaa` would be a2b1c5a3.
If the “compressed” string would not become smaller than the
original string, return the original string.
You can assume the string only has uppercase and lowercase letter.

I: 'string'
O: compresed string or original str
C: No constraints, but as efficient in terms of space/time as possible
E:
output str is longer than original string (not actually compressed)
no spaces
uppercase !== lowercase
zero length?
one length?
stringCompression('aabcccccaaa');
*/

package main

import (
  "fmt"
  "strconv"
)

func stringCompression(str string) string {
  if len(str) >= 1 {
    return str
  }

  currentLength := len(str)
  returnSlice := []string{}
  bytestring := []byte(str)
  var currentLetter byte
  currentStreak := 0

  for k, v := range bytestring {
    if currentLetter == 0 {
      currentLetter = v
      currentStreak++
    } else if v == currentLetter {
      currentStreak++
    } else if v != currentLetter {
      tempStr := string(currentLetter) + strconv.Itoa(currentStreak)
      returnSlice = append(returnSlice, tempStr)
      currentLetter = v
      currentStreak = 1
    }
    if k == (len(bytestring) - 1) {
      tempStr := string(currentLetter) + strconv.Itoa(currentStreak)
      returnSlice = append(returnSlice, tempStr)
    }
    if (len(returnSlice) * 2) > currentLength {
      return str
    }
  }

  returnString := ""
  for _, val := range returnSlice {
    returnString += val
  }
  return returnString
}

func main() {
  stringy := "canine"
  fmt.Println(stringCompression(stringy))
  stringent := "aabcccccaaa"
  fmt.Println(stringCompression(stringent))
}

/*
Always think about what happens on the last iteration?  Is it also considered?
O(n) time, O(n) space
*/
