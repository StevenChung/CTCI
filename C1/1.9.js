/*
1.9) String Rotation
Assume you have a method `isSubstring` which checks if one word is a substring of another.
  Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using
  only a call to `isSubstring` (“waterbottle” is a rotation of `erbottlewat`)
  waterbottlewaterbottle
*/

var isSubstring = (s1, s2) => {
  if (s1.length === s2.length) {
    return s1 === s2;
  }

  let longerStr = (s1.length > s2.length) ? s1 : s2;
  let shorterStr = (s1.length < s2.length) ? s1 : s2;

  for (let i = 0; i <= longerStr.length - shorterStr.length; i++) {
    let fragment = longerStr.slice(i, shorterStr.length);
    if (shorterStr === fragment) {
      return true;
    }
  }

  return false;
};

isSubstring(s1 + s1, s2);
// Only one call to isSubstring to see if s2 is a rotation of s1
// it must be the case that, if s2 is a rotation, that it would appear in two iterations of s1
// waterbottlewaterbottle must contain erbottlewat (if erbottlewat is a substring)
