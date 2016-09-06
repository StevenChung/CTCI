/*
1.2) Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
I: Two Strings
O: Boolean
C: None
E:
if same length, check to see if they're the same (naive first check)
to be a permutation, character counts would be the same

What if they're different length?
What constitutes a permutation of the other?
*/

var checkPermutation = (str1, str2) => {
  if (str1.length !== str2.length) {
    return false;
  }
  // naive solution
  // two objects to track character counts of each
  // requires iterating through each object once (O(n) * 2 (assuming same length strengths because of above))

  // finally, requires iterating through both objects to see if counts are the same
  // O(n) solution
  let obj1 = {};
  let obj2 = {};
  for (let i = 0; i < str1.length; i++) {
    obj1[str1[i]] = (str1[i] || 0) + 1;
    obj2[str2[i]] = (str2[i] || 0) + 1;
  }
  for (let key in obj1) {
    if (!(key in obj2)) {
      return false;
    }
    if (obj1[key] !== obj2[key]) {
      return false;
    }
  }
  return true;
};
// O(n) space // O(n) time (for loops with O(1) operations)
// for string questions, always confirm:
/*
a) space?
b) case sensitivity?
c) what other edge cases?

Two solutions: sort the two strings and compare letters
O(n log n) for time (sorting), but it's cleaner

Above solution is O(n) (two-scan) approach
