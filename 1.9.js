/*
1.9) String Rotation
Assume you have a method `isSubstring` which checks if one word is a substring of another.
  Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using
  only a call to `isSubstring` (“waterbottle” is a rotation of `erbottlewat`)
*/

var isSubstring = (s1, s2) => {
  if (s1.length === s2.length) {
    return (s1 === s2);
    // if they're the same length, there's no "sub" feasibility, so they're either equal or they're not
  } else {
    let longerStr = (s1.length > s2.length) ? s1 : s2;
    let shorterStr = (s1.length < s2.length) ? s1 : s2;
    let diff = longerStr.length - shorterStr.length;
    // a substring can only exist from a certain index on
    // the shorter substring can only exist from index 0 to index diff (difference in
    // both strings' lengths)
    // i.e. 'oil' can, at the latest, be at index 1 for 'foil' (4 - 3)
    for (let i = 0; i <= diff; i++) {
      if (longerStr[i] === shorterStr[0]) {
        // if given letter matches first letter of shorterStr
        if (longerStr.slice(i, i + shorterStr.length) === shorterStr) {
          // for substrings, do not forget slice
          // you had a dumb approach before!
          // slice longerStr from current index for the length of the shorter str
          // if they're equal, it's a substring
          return true;
        }
      }
    }

    return false;
  }
};

var isRotation = (s1, s2) => {
  // problem calls for "a call" to subString, but there is iteration here
  // this is v0, so there is probably some optimization to be had

  if (s1 === s2) {
    return true;
  }
  for (let i = 1; i < s1.length; i++) {
    // if s1 sliced to the end is a substring of s2
    if (isSubstring(s1.slice(i, s1.length), s2)) {
      if ((s1.slice(i, s1.length) + s1.slice(0, i)) === s2) {
        // check to see if that + the initial part of s1 is equal to s2
        // it's the slice from i to s1.length first because
        // the inverse would just be s1
        // also, it's the sum of only two slices because any other permutation would be caught at a previous index
        /*
        'waterbottle', 'terbottlewa'
        'aterbottlew' does not need to be checked on terbottle (index 2) because
        'aterbottlew' would have been checked at the previous index
        */
        return true;
      }
    }
    if (isSubstring(s2.slice(i, s2.length), s1)) {
      // checking s2 over s1 now for the same as above
      if ((s2.slice(i, s2.length) + s2.slice(0, i)) === s1) {
        return true;
      }
    }
  }
  return false;
};
