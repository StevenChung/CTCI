/*
Implement BubbleSort
[10, 50, 72, 1, 3, 8, -13, 20]
*/

const bubbleSort = (arr) => {
  let endIndex = arr.length - 1;
  while (endIndex > 0) {
    let swapDone = false;
    for (let i = 0; i < endIndex; i++) {
      if (arr[i] > arr[i + 1]) {
        let tempOld = arr[i];
        arr[i] = arr[i + 1];
        arr[i + 1] = tempOld;
        swapDone = true;
      }
    }
    if (!swapDone) {
      break;
    }
    endIndex--;
  }
  return arr;
};

console.log(bubbleSort([10, 50, 72, 1, 3, 8, -13, 20]));
