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
  [5, 10, 17],
  [3, 9, -1],
  [2, 4, 12]
]);


rotateMatrix([
  [1,2,3,4,5],
  [6,7,8,9,10],
  [11,12,13,14,15],
  [16,17,18,19,20],
  [21,22,23,24,25],
], true);
*/

var rotateMatrix = (matrix, clockwise = true) => {
  let returnMatrix = matrix.slice();
  let even = ((matrix[0].length % 2) === 0) ? true : false;
  let iterationCount = Math.floor(matrix[0].length / 2);
  let min = 0;
  let max = matrix[0].length;

  if (clockwise) {
    while (iterationCount > 0) {
      let firstRow = matrix[min].slice(min, max);
      let lastRow = matrix[max - 1].slice(min, max);
      let firstColumn = [];
      let lastColumn = [];
      for (let i = min; i < max; i++) {
        firstColumn.push(matrix[i][min]);
        lastColumn.push(matrix[i][max - 1]);
      }

      for (let j = max - 1, i = 0; j >= min; j--, i++) {
        returnMatrix[max - 1][j] = lastColumn[i];
        returnMatrix[min][j] = firstColumn[i];
      }

      for (let i = min, j = 0; i < max; i++, j++) {
        returnMatrix[i][max - 1] = firstRow[j];
        returnMatrix[i][min] = lastRow[j];
      }

      min++;
      max--;
      // first row becomes last column
        // maintain order
      // last column becomes last row DONE
        // reverse order
      // last row becomes first column
        // maintain order
      // first column becomes first row DONE
        // reverse order
      iterationCount--;
    }
  } else {
    while (iterationCount > 0) {
      let firstRow = matrix[min].slice(min, max);
      let lastRow = matrix[max - 1].slice(min, max);
      let firstColumn = [];
      let lastColumn = [];
      for (let i = min; i < max; i++) {
        firstColumn.push(matrix[i][min]);
        lastColumn.push(matrix[i][max - 1]);
      }

      for (let j = max - 1, i = 0; j >= min; j--, i++) {
        returnMatrix[j][min] = firstRow[i];
        returnMatrix[j][max - 1] = lastRow[i];
      }

      for (let i = min, j = 0; i < max; i++, j++) {
        returnMatrix[max - 1][i] = firstColumn[j];
        returnMatrix[min][i] = lastColumn[j];
      }

      min++;
      max--;
      // first row becomes first column
        // reverse order
      // last column becomes first row DONE
        // maintain order
      // last row becomes last column
        // reverse order
      // first column becomes last row DONE
        // maintain order
      iterationCount--;
    }
  }

  return returnMatrix;
};
