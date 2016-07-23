/*

1.6) String Compression
Implement a method to perform a basic compression using the counts of
repeated characters.  For example, the string `aabcccccaaa` would be a2b1c5a3.
If the “compressed” string would not become smaller than the
original string, return the original string.
You can assume the string only has uppercase and lowercase letter.

I: 'string'
O: compresed string or original str
C: No constraints, but as efficient in terms of space/time as possible
E:
output str is longer than original string (not actually compressed)
no spaces
uppercase !== lowercase
zero length?
one length?
stringCompression('aabcccccaaa');
*/
var stringCompression = (str) => {
  let compStr = '';

  let curLetter = str[0];
  let trackCount = 1;
  // seeded for the first letter/count
  for (let i = 1; i < str.length; i++) {
    if (curLetter === str[i]) {
      trackCount++;
    } else {
      compStr += curLetter + trackCount; // append to compStr
      curLetter = str[i]; // reset
      trackCount = 1;
    }
    if (i === str.length - 1) {
      compStr += curLetter + trackCount;
      // on the last, we need to re-append to compStr
      // otherwise, it's left off because it never hits the above else
    }
  }
  return (compStr.length < str.length) ? compStr : str;
  // only return if it is shorter, not equal, not larger
  // O(n) time
  // O(n) space (return Str is slowly built)
};
