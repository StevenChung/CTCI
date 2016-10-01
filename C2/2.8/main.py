# https://leetcode.com/problems/linked-list-cycle-ii/


# I: Linked List
# O: Node of cycle or null
# C: O(1)
# E:
# 1 -> 2 (no cycle)
# 1 -> 2 -> 3 -> 4 -> 2 (cycle at non-head of linkedlist)
# 1 -> 2 -> 3 -> 1 (cycle at head of linked list)


class ListNode(object):
    def __init__(self, val):
        self.val = val
        self.next = None


def detectCycle(head):
    slow, fast = head, head
    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next.next

        if id(slow) == id(fast):
            slow = head

# 1) Detect if there's a loop
# - Standard
# 2) If they do meet, move slow to the start
# 3) Increment both by one only until they meet
#
# l = length of the cycle
# m = distance from head to start of cycle
# k  = distance of the meeting point of S/F from the start of the loop
# DistanceOfSlow = m + (p * l) + k
# (p is some number of cycles since we don't knwo hoa mny until they met)
# DistanceOfFast = m + (q * l) + k
# (q is some number of cycles since we don't knwo hoa mny until they met)
# Since we know slow moves only half as fast, it must be the case that...
#
#
# m + (q * l) + k = 2(m + (p * l) + k)
# reduces to
# m + k = (l) * (q - 2p)
# this means that m + k is a multiple of the length of the cycle
# k is where F is at (since it is where it met S)
# moving S to the start means that it will move m to the start
# F will travel m + (having started at k) k
# Thus, they must meet the start since m + k is multiple of l























# big circle where cycle starts at head
bigCycleList = ListNode(5)
bigCycleList.next = ListNode(4)
bigCycleList.next.next = ListNode(3)
bigCycleList.next.next.next = bigCycleList
print(detectCycle(bigCycleList).val)

# # cycle starts at non-head
cycleList = ListNode(5)
cycleList.next = ListNode(4)
cycleList.next.next = ListNode(3)
cycleList.next.next.next = ListNode(2)
cycleList.next.next.next.next = ListNode(1)
cycleList.next.next.next.next.next = cycleList.next.next
print(detectCycle(cycleList).val)

#non-cycle linkedlist
nonCycleList = ListNode(5)
nonCycleList.next = ListNode(4)
nonCycleList.next.next = ListNode(3)
nonCycleList.next.next.next = ListNode(2)
print(detectCycle(nonCycleList))

print(detectCycle(None))
