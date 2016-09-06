/*
1.7) Rotate Matrix
Given an image represented by an NxN matrix,
 where each pixel in the image is represented by 4 bytes,
  write a method to rotate the image by 90 degrees.
Can you do this in place? (i.e. modifiying in place)
*/

package main

import (
  "fmt"
)

func RotateMatrix (matr [][]int, clockwise bool) [][]int {
  var returnMatrix = [][]int{}
  if clockwise == true {
    for col := 0; col < len(matr); col++ {
      tempSlice := []int{}
      for row := len(matr) - 1; row >= 0; row-- {
        tempSlice = append(tempSlice, matr[row][col])
      }
      returnMatrix = append(returnMatrix, tempSlice)
    }
  } else {
    for col := len(matr) - 1; col >= 0; col-- {
      tempSlice := []int{}
      for row := 0; row < len(matr); row++ {
        tempSlice = append(tempSlice, matr[row][col])
      }
      returnMatrix = append(returnMatrix, tempSlice)
    }
  }
  return returnMatrix
}

func main() {
  matrix := [][]int{
    []int{1,2,3},
    []int{4,5,6},
    []int{7,8,9},
  }
  matrixb := RotateMatrix(matrix, true)
  fmt.Println(matrixb)
  fmt.Println(matrix)
}
