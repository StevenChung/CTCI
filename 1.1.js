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

/*
Hints from the back:
1) Try a hash table (done)
2) Can you achieve O(n log n) time? (done? Isn't the above just (O(n)) on time?)
If you specifically wanted n log n, you could sort the str and check to see if adjacent are equal.
`n log n` is an indication of sorting!
3) Could a bit vector be useful? https://en.wikipedia.org/wiki/Bit_array

Ask about assumptions off the bat
If str.length > 128 (ASCII set), it cannot be, so we could save some modicum of work

For no additional structures, you would have to check every previous letter before it
this ensures O(1), but O(n^2) time.
