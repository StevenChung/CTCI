/*
1.5) One Away: There are three types of edits to a given string: insert a character, remove a character, or replace a character.
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
  // the diff between the lengths needs to be between (inclusive) -1 and 1

  if (str1.length === str2.length) {
    if (str1 === str2) {
      return true;
    }

    let booboo = false;
    for (let i = 0; i < str1.length; i++) {
      if (str1[i] !== str2[i]) {
        if (!booboo) {
          booboo = true;
        } else {
          return false;
        }
      }
    }
  }

  let longerString = (str1.length > str2.length) ? str1 : str2;
  let shorterString = (str1.length > str2.length) ? str2 : str1;
  let booboo = false;
  for (let i = 0; i < longerString.length; i++) {
    if (!booboo) {
      if (longerString[i] !== shorterString[i]) {
        booboo = true;
        // ple pale
      }
    } else {
      if ((longerString[i] !== shorterString[i - 1]) && (longerString[i] !== shorterString[i])) {
        return false;
      }
    }
  }

  return true;
};


/// WRITE ALL PERMUTATIONS / CONSIDER ALL PERMUTATIONS
// DO NOT JUMP INTO CODING
// PSEUDOCODE
// WRITE DIAGRAMS ETC.
// BE PATIENT
