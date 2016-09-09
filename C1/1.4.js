/*
1.4) Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
A palindrome is a word or phrase that is the same forwards and backwards.
A permutation is a rearrangement of letters.
The palindrome does not need to be limited to just dictionary words.

Example: ‘Tact Coa’
Output: True (Taco cat)

I: String
O: boolean
C: As efficient as possible in terms of space and time
E: Length is 1 or less => Palindrome

There can must be one character that has an odd count and only if the overall length is odd;
if length is even, there can be no count of 1

Base case is character length, then deciding how to proceed
always check for errors early on
ALWAYS THINK ABOUT SPACES ALWAYS ASK ABOUT SPACES
ASK ABOUT CASE SENSITIVITY
Gail recommends: "What are defining features of a palindrome?"
*/


var palindromePermutation = (str) => {
  if (typeof str !== 'string') {
    return false;
  }

  if (str.length <= 1) {
    return true;
  }

  let obj = {};
  let charLength = 0;

  for (let i = 0; i < str.length; i++) {
    if (str.charAt(i) !== ' ') {
      obj[str.charAt(i)] = (obj[str.charAt(i)] || 0) + 1;
      charLength++;
    }
  }

  let even = ((charLength % 2) === 0) ? true : false;
  let oddCountAlready = false;

  for (let key in obj) {
    if (even) {
      if ((obj[key] % 2) !== 0) {
        return false;
      }
    } else {
      if ((obj[key] % 2) !== 0) {
        if (oddCountAlready) {
          return false;
        } else {
          oddCountAlready = true;
        }
      }
    }
  }

  return true;
  // use object to keep character counts
  // if even, all must be even counts
  // if odd, only one character can have an odd count
};

// Your original (above) is O(n), length of string
