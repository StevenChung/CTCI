/*
1.7) Rotate Matrix
Given an image represented by an NxN matrix, where each pixel in the image is represented by 4 bytes, write a method to rotate the image by 90 degrees.
  Can you do this in place? (i.e. modifiying in place)

I: NxN array (square)
O: NxN Array
C: None on space/time, but optimized (obv)
E:
Which direction? Clockwise? Counterclockwise?
1x1?
2x2?
3x3...?
Even/odd will have different because center of odd does not change, but
there is no center in even

rotateMatrix([
  [5, 10],
  [3, 9]
]);

rotateMatrix([
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
]);

[
[7, 4, 1],
[8, 5, 2],
[9, 6, 3]
]

rotateMatrix([
  [1,2,3,4,5],
  [6,7,8,9,10],
  [11,12,13,14,15],
  [16,17,18,19,20],
  [21,22,23,24,25],
]);
*/

var rotateMatrix = (matrix, clockwise = true) => {
  let returnMatrix = [];
  if (clockwise) {
    for (let i = 0; i < matrix.length; i++) {
      let tempArr = [];
      for (let j = matrix.length - 1; j >= 0; j--) {
        tempArr.push(matrix[j][i]);
      }
      returnMatrix.push(tempArr);
    }
  } else {
    for (let i = matrix.length - 1; i >= 0; i--) {
      let tempArr = [];
      for (let j = 0; j < matrix.length; j++) {
        tempArr.push(matrix[j][i]);
      }
      returnMatrix.push(tempArr);
    }
  }
  return returnMatrix;
};
/*
diagrams are very important
try for smaller versions first
for clockwise, the first column becomes the first row, but in reverse order
second column becomes the second row, but reverse order
for clockwise, the last column becomes the first row, but in the same order

NOTES:
For both JS and Go, Array.prototype.slice and copy() work similarly
It'll copy primitive values, so...
var myArr = [1,2,3]
var newArr = myArr.slice()
newArr[0] = 'canine'
// newArr = ['canine', 2, 3]
// myArr = [1,2,3]

HOWEVER, consider:
var myArr = [
[0, 1],
[2, 3],
];
var newArr = myArr.slice();

newArr[0][0] = 'dog'
This will also mutate myArr because slice copies the primitive values, but
arrays (arrays and slices in Go) are not primitive.
Thus, slice and copy does not generate a new array/slice without reference
to the original.

newArr[0] = [99, 99]
This does not mutate myArr because this is simply mutating the outer array.
(changing the array that element points at)

newArr[0][0] = 'dog'
This, on the other hand, changes an element in an original inner slice

Edit: tldr: slices in go are like pointers to arrays. Here
we are dealing with a
pointer to an array of pointers.

Thus, the first example, you are changing the pointer to a completely
different array.  In the second example, you are changing the pointer
inside the outer pointer (which happens to be the original)
*/
