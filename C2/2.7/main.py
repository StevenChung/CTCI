# Intersection : Given  two  (singly) linked lists,
#  determine  if  the  two  lists intersect. Return the inter-
# secting node. Note that the intersection  is  defined
# based on reference,  not  value. That  is,  if  the  kth
# node  of  the first linked list  is  the exact  same
# node (by reference)  as  the  jth  node  of  the second
# linked list, then they  are  intersecting.

class ListNode(object):
	def __init__(self, val):
		self.val = val
		self.next = None

# two pointers
def getIntersectionNode(headA, headB):
	if headA == None or headB == None:
		return None

	pointerOne = headA
	pointerTwo = headB

	while pointerOne != pointerTwo:

        if pointerOne == None:
            pointerOne = headB
        else:
            pointerOne = pointerOne.next

        if pointerTwo == None:
            pointerTwo = headA
        else:
            pointerTwo = pointerTwo.next


	return pointerOne

# If pointer one reaches the end of its list before pointer two, switch
# it to point to the other list; vice versa also applies.
#
# Basically, think of it as if you combined the two lists.
#
# If one list is length 4 and the other is list 6 and there are no intersections,
# the exit condition is if both are None at the same time, which would happen regardless of
# length disparity.  If they do intersect and have the same number of nodes before the intersection,
# that's a simple catch since they're both iterating at the same rate.
#
# If they're different, the longer difference between the two will be accounted for in the second pass
#
# The other assumption is that once the lists intersect, they will finish at the same place/nodes
#
# That is to say, they will not intersect once then deviate (or so leetcode believes)...
#
# However, this is solved with the hashtable method below.





# hash table
# reminder: python also has sets (just values in an unordered data structure that
# can easily check membership with "in"
def getIntersectionNode(headA, headB):
	if headA == None or headB == None:
		return None

	hashTable = {}
	pointerA = headA
	while pointerA != None:
		if pointerA.val in hashTable:
			hashTable[pointerA.val].append(pointerA)
		else:
			hashTable[pointerA.val] = [pointerA]
		pointerA = pointerA.next
	print(hashTable)
	pointerB = headB
	while pointerB != None:
		if pointerB.val in hashTable:
			for _, v in hashTable[pointerB.val]:
				if v == pointerB:
					return pointerB

	return None
