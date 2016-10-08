# For the purposes of this question, a balanced tree is defined to be a tree
# Implement a function to check if a binary tree is balanced.
# such that the heights of the two subtrees of any node
# never differ by more than one
#
# https://leetcode.com/problems/balanced-binary-tree/

class TreeNode(object):
	def __init__(self, val):
		self.val = val
		self.left = None
		self.right = None


def isBalanced(root):
	def checkNode(node):
		if node is None:
			return 0

		left = checkNode(node.left)
		if left != -1:
			right = checkNode(node.right)
			if right != -1:
				if abs(right - left) > 1:
					return -1
				else:
					return max(left, right) + 1
			else:
				return -1
		else:
			return -1
	return checkNode(root) != -1
