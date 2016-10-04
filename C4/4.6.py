# 4.6 Successor: Write an algorithm to find the "next" node (in-order successor)
# of a given node in a binary search tree.  Each node has a link to its parent
# https://leetcode.com/problems/inorder-successor-in-bst/



class TreeNode(object):
	def __init__(self, val):
		self.val = val
		self.left = None
		self.right = None

##py3 has nonlocal
def inorderSuccessor(root, nodeToCheck):
	# in order traversal from root until you hit p
	# then whatever is next, return that? or null
	arrOfNodes = []
	indexOfP = None

	def inOrderCheck(node):
		nonlocal indexOfP
		if node != None:
			inOrderCheck(node.left)
			arrOfNodes.append(node)
			if node == nodeToCheck:
				indexOfP = len(arrOfNodes) - 1
			inOrderCheck(node.right)

	inOrderCheck(root)
	if indexOfP is None or (indexOfP + 1) == len(arrOfNodes):
		return None
	else:
		return arrOfNodes[indexOfP + 1]


## py2 does not have nonlocal, so dump it in a dict
def inorderSuccessor(root, nodeToCheck):
	# in order traversal from root until you hit p
	# then whatever is next, return that? or null
	hashTable = {
	  'arrOfNodes': [],
	  'indexOfP': None
	}

	def inOrderCheck(node):
		if node != None:
			inOrderCheck(node.left)
			hashTable['arrOfNodes'].append(node)
			if node == nodeToCheck:
				hashTable['indexOfP'] = len(hashTable['arrOfNodes']) - 1
			inOrderCheck(node.right)

	inOrderCheck(root)
	if hashTable['indexOfP'] is None or (hashTable['indexOfP'] + 1) == len(hashTable['arrOfNodes']):
		return None
	else:
		return hashTable['arrOfNodes'][hashTable['indexOfP'] + 1]








# Testing
testTree = TreeNode(1)
testTree.left = TreeNode(2)
testTree.right = TreeNode(3)
testTree.left.left = TreeNode(4)
testTree.left.right = TreeNode(5)
print(inorderSuccessor(testTree, testTree.left.left).val)
