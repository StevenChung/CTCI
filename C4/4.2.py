# Minimal Tree: Given a sorted (increasing order) array
# with unique integer elements, write an algorithm
#  to create a binary search tree with minimal height.
# https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
# I: array
# O: binary search tree
# C: None
# E: odd length arrays
# even length arrays
# empty arrays
# length of one

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def sortedArrayToBST(self, nums):
        """
        :type nums: List[int]
        :rtype: TreeNode
        """
        if len(nums) == 0:
            return
        if len(nums) == 1:
            node = TreeNode(nums[0])
            return node
        if len(nums) == 2:
            node = TreeNode(nums[0])
            node.right = TreeNode(nums[1])
            return node

        if len(nums) % 2 == 0:
            middle = int((len(nums) / 2) - 1)
        else:
            middle = int(len(nums) / 2)

        node = TreeNode(nums[middle])
        node.left = self.sortedArrayToBST(nums[0:middle])
        node.right = self.sortedArrayToBST(nums[middle+1:len(nums)])
        return node
