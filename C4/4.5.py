# Validate BST: Implement a function to check if a binary tree is a binary
# search tree
# Binary Search Tree is one, for a given node n, where nodes to the left are
# less than or equal to n and nodes to the right are greater than or equal

# Basic Case first
# The naive solution is to check every node that is a child of node.left
# to ensure that it is less than node
# Check every node that is a child of node.right to ensure that it is greater
#
# Duplicate work in that we check all subsequent nodes frequently
#
# Instead, check to see that for a given node(n), if it has a left(nl), that nl.r
# is greater than nl (basic check) and that it is less than n
#
# Doing this search on all subsequent nodes ensures that duplicate work is reduced
# Also, enforcement of this rule ensures that we fulfill the ruleset from above
#
# Think of duplicate work/checks as are there any overlaps in this ruleset?


# I
#     Confirm that input is always a binary tree
# O
#     boolean
# C
#     none
# E

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# for a binary search tree, an in-order traversal should be a sorted array
# this attempt is recursive
# RECURSIVE
class Solution(object):
    def isValidBST(self, root):
        purportedlySortedArr = []

        def inOrderPush(node):
            if node != None:
                inOrderPush(node.left)
                purportedlySortedArr.append(node.val)
                inOrderPush(node.right)
            return

        inOrderPush(root)

        if len(purportedlySortedArr) > 0:
            for i in range(0, len(purportedlySortedArr)):
                if i > 0 and purportedlySortedArr[i] <= purportedlySortedArr[i - 1]:
                    return False
        return True


## TEST SUITE
bstTest = BinaryTreeNode(100)
bstTest.left = BinaryTreeNode(25)
bstTest.right = BinaryTreeNode(175)
bstTest.left.left = BinaryTreeNode(6)
bstTest.left.right = BinaryTreeNode(38)
print(isValidBST(bstTest))
