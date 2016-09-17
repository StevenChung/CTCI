var Node = function (val) {
  this.val = val;
  this.next = null;
};

class Queue {
  // FIFO
  constructor() {
    this.last = null;
    this.first = null;
  }

  add(val) {
    let newNode = new Node(val);
    if ((this.first === null) && (this.last === null)) {
      this.first = newNode;
      this.last = newNode;
    } else {
      let oldLast = this.last;
      oldLast.next = newNode;
      this.last = newNode;
    }
  }

  remove() {
    if ((this.first === null) && (this.last === null)) {
      return null;
    } else {
      var nodeToReturn = this.first;
      if (nodeToReturn.next === null) {
        this.first = null;
        this.last = null;
      } else {
        this.first = nodeToReturn.next;
      }

      return nodeToReturn.val;
    }
  }

  peek() {
    return this.first.val;
  }

  isEmpty() {
    if (this.first === null) {
      return true;
    } else {
      return false;
    }
  }
}

/*
var q = new Queue();
q.add(5);
q.add(7);
q.add(8);
*/
