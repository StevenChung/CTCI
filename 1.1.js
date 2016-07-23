/*
1.1) Is Unique? Implement an algorithm to determine if a string has all unique characters.
  What if you cannot use additional data structures?



I: string
O: boolean
C: no additional data structures (O(1) space)
E: string is just one character

*/

var isUnique = (str) => {
  if (str.length <= 1) {
    return true;
  }
  // use of a hash table to see if a given key has been assigned already, but that takes more than O(1)
  // also, that involves an additional data structure
  let obj = {};
  for (let i = 0; i < str.length; i++) {
    if (str[i] in obj) {
      return false;
    } else {
      obj[str[i]] = str[i];
    }
  }
  return true;
};
