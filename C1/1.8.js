// 1.8) Zero Matrix
// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column are set to 0.
/*
I: M x N Matrix (Array nested in an array)
O: M x N Matrix with zeroes where there was previously a Zero
C: There are no space/time constraints stated.  However, efficiency is valued.
E:
Mx1 / 1xN matrix?
What if there is more than one zero in a given row/index?  Means that we probably
do not want to convert all to zeroes until the end.
will mutate original matrix?

zeroMatrix([
  [0, 3, 5],
  [1, 2, 0],
  [7, 3, 1],
  [1, 3, 9]
]);


*/
var zeroMatrix = (matrix) => {
  let copyMat = matrix.slice();
  // iterate through original matrix
  let obj = {
    row: {},
    col: {},
  };
  // obj instead of array to de-dupe and not perform extra rows/cols
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      // if zero, push [r][c] to object to track / de-dupe
      if (matrix[i][j] === 0) {
        obj.row[i] = i;
        obj.col[j] = j;
      }
    }
  }

  // iterate through obj.row / obj.col
  for (let key in obj.row) {
    for (let i = 0; i < matrix[key].length; i++) {
      // convert each row/col to zeroes
      copyMat[key][i] = 0;
    }
  }

  for (let key in obj.col) {
    for (let i = 0; i < matrix.length; i++) {
      copyMat[i][key] = 0;
    }
  }
  // return matrix
  return copyMat;
};
