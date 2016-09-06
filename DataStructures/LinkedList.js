var LinkedList = function () {
  this.head = null;
  this.tail = null;
  this.length = 0;
};

LinkedList.prototype.addNode = function (value) {
  return {
    value: value,
    next: null,
  };
};

LinkedList.prototype.addToTail = function (value) {
  this.length++;
  var newNode = this.addNode(value);
  if (this.head === null) {
    this.head = newNode;
    this.tail = newNode;
    return;
  }

  this.tail.next = newNode;
  this.tail = newNode;
};

LinkedList.prototype.removeHead = function() {
  this.head = this.head.next;
  this.length--;
};

LinkedList.prototype.removeTail = function() {
  if (this.length === 1) {
    this.head = null;
    this.tail = null;
    this.length--;
  } else {
    var currentNode = this.head;
    var distance = this.length - 2;
    for (var i = 0; i < distance; i++) {
      currentNode = currentNode.next;
    }
    currentNode.next = null;
    currentNode = this.tail;
    this.length--;
  }
};

LinkedList.prototype.addToHead = function (value) {
  this.length++;
  if (this.head === null) {
    this.head = this.addNode(value);
    return;
  } else {
    var oldHead = this.head;
    this.head = this.addNode(value);
    this.head.next = oldHead;
    return;
  }
};

LinkedList.prototype.contains = function (searchQuery) {
  if ((this.head.value === searchQuery) || (this.tail.value === searchQuery)) {
    return true;
  }
  // the above is a catch for if there's only one item in the LinkedList

  var currentNode = this.head;
  while (currentNode.next !== null) {
    if (currentNode.value === searchQuery) {
      return true;
    }
    currentNode = currentNode.next;
  }

  return false;
};


/*
var newList = new LinkedList();
newList.addToTail(7);
newList.addToTail(8);
newList.addToTail(15);
newList.addToHead(22);
newList.addToHead(14);
newList.removeHead();
newList; // 22, 7, 8, 15
newList.removeTail();
newList; // 22, 7, 8
*/

/*
Linked List
Singly Linked: Node has a next property
Doubly Linked: Node has a next and previous property

Unlike an array, Linked List does not provide index.
This means there is not constant time lookup for
a particular index (aside from head/tail)

Thus, to find the Kth element, you would need to
iterate

Inserting into an array generally takes O(n) because
... inserting at the end is O(1),
but anywhere else requires
changing indices of other elements (thus O(n))
However, we have constant time lookup for
numerical indices!

In Single Linked List
Adding head is constant
Removing head is constant
Removing Tail is linear (O(n)) because
we must iterate to the next to last (SINGLE LINKED)
Adding Tail is constant
*/
