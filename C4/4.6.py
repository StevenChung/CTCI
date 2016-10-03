# 4.6 Successor: Write an algorithm to find the "next" node (in-order successor)
# of a given node in a binary search tree.  Each node has a link to its parent
# https://leetcode.com/problems/inorder-successor-in-bst/



class TreeNode(object):
	def __init__(self, val):
		self.val = val
		self.left = None
		self.right = None

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
	print(arrOfNodes)
	print(indexOfP, 'why not zero here?')












# Testing
testTree = TreeNode(1)
testTree.left = TreeNode(2)
testTree.right = TreeNode(3)
testTree.left.left = TreeNode(4)
testTree.left.right = TreeNode(5)
inorderSuccessor(testTree, testTree.left.left)
print(testTree.left.left == testTree.left.left)
