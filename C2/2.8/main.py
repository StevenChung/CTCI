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
            fast = head

1) Detect if there's a loop
- Standard
2)


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
