/* 1.3) URLify: Replace all spaces in a string with ‘%20’.
I: String
O: String with spaces replaced
C: Highest Efficiency
E: Ends in space, starts with space, just a space, consecutive spaces?
*/

var URLify = (str) => {
  // naive
  // iterate through string
    // not just pure split
    // push everything before that is not a string to array, then push a '%20'
  // at the end, join array and return that str
  // O(n) time, O(n) space
  // could you do it in O(1) space?
  return str.split('').map((char) => (char === ' ') ? '%20' : char).join('');
};

// seems trivial in JS with the above
