/*
1.9) String Rotation
Assume you have a method `isSubstring` which checks if one word is a substring of another.
  Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using
  only a call to `isSubstring` (“waterbottle” is a rotation of `erbottlewat`)
*/

var isSubstring = (s1, s2) => {
  if (s1.length === s2.length) {
    return (s1 === s2);
  } else {
    let longerStr = (s1.length > s2.length) ? s1 : s2;
    let shorterStr = (s1.length < s2.length) ? s1 : s2;
    let diff = longerStr.length - shorterStr.length;

    for (let i = 0; i <= diff; i++) {
      if (longerStr[i] === shorterStr[0]) {
        if (longerStr.slice(i, i + shorterStr.length) === shorterStr) {
          // for substrings, do not forget slice
          // you had a dumb approach before!
          return true;
        }
      }
    }

    return false;
  }
};

var isRotation = (s1, s2) => {
  if (s1 === s2) {
    return true;
  }
  for (let i = 1; i < s1.length; i++) {
    if (isSubstring(s1.slice(i, s1.length), s2)) {
      if ((s1.slice(i, s1.length) + s1.slice(0, i)) === s2) {
        return true;
      }
    }
    if (isSubstring(s2.slice(i, s2.length), s1)) {
      if ((s2.slice(i, s2.length) + s2.slice(0, i)) === s1) {
        return true;
      }
    }
  }
  return false;
};
