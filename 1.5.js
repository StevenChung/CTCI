/*
1.5) One Away: There are three types of edits to a given string: insert a character,
 remove a character, or replace a character.
Given two strings, write a function to check if they are one edit away (or zero).

I: Two strings
O: Boolean
C: Time/Space => None.  However, it doesn't hurt to offer the optimal solution.
E:

Same String
Strings of different length
What if a string is length 1? Length 0?
I.e. what if they're both the same length (for 1 or 0)?
Is it one edit per string or one total edit?
*/

var oneAway = (str1, str2) => {
  if (str1 === str2) {
    return true;
  } else if (Math.abs(str1.length - str2.length) > 1) {
    return false;
  } else {
    var hashitout = (str) => str.split('').reduce((a, b) => {
      a[b] = (a[b] || 0) + 1;
      return a;
    }, {});
    var hashOne = hashitout(str1);
    var hashTwo = hashitout(str2);
    var booboo = 0;

    for (var key in hashTwo) {
      if (!(key in hashOne) || (hashOne[key] !== hashTwo[key])) {
        booboo++;
      }

      if (booboo > 1) {
        return false;
      }
    }

    return true;
  }
};

/// WRITE ALL PERMUTATIONS / CONSIDER ALL PERMUTATIONS
// DO NOT JUMP INTO CODING
// PSEUDOCODE
// WRITE DIAGRAMS ETC.
// BE PATIENT
// consider all possibilities/edge cases
// i would try to not erase IOCE, but other stuff should be safe
