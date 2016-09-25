# 4.3 Given a binary tree, design an algorithm which creates a linked list of
# all nodes at each depth

class BinaryTreeNode():
	def __init__(self, val):
		self.Value = val
		self.Left = None
		self.Right = None

class LinkedListNode():
	def __init__(self, val):
		self.Value = val
		self.Next = None




rooty = BinaryTreeNode(1)
rooty.Left = BinaryTreeNode(2)
rooty.Right = BinaryTreeNode(3)
rooty.Left.Left = BinaryTreeNode(4)
rooty.Left.Right = BinaryTreeNode(5)
rooty.Right.Left = BinaryTreeNode(6)
rooty.Right.Right = BinaryTreeNode(7)

def LinkedListAtDepths(rootNode):
	# create queue that initially has just the rootnode
	# we use BF iteration to append items onto a list
	# create linkedList from those items
	# also have a temporary array that has the children from that BF iteration
	# reset the queue to the array of children nodes
	queue = [rootNode]
	returnArrOfLinkedListNodes = []
	while len(queue) > 0:
		nextQueue = []
		llHead = None
		llNode = None
		while len(queue) > 0:
			item = queue.pop(0)
			if item.Left != None:
				nextQueue.append(item.Left)
			if item.Right != None:
				nextQueue.append(item.Right)
			if llHead == None:
				llNode = LinkedListNode(item.Value)
				llHead = llNode
			else:
				llNode.Next = LinkedListNode(item.Value)
				llNode = llNode.Next

		returnArrOfLinkedListNodes.append(llHead)
		if len(nextQueue) > 0:
			queue = nextQueue

	return returnArrOfLinkedListNodes


testArr = LinkedListAtDepths(rooty)
for i in range(len(testArr)):
	node = testArr[i]
	while node != None:
		print(node.Value)
		node = node.Next
