package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    dummyHead := &ListNode{}
    current := dummyHead

    if list1.Val < list2.Val {
		current = list1;
		dummyHead = list1;
		list1 = list1.Next;
	} else {
		current = list2;
		dummyHead = list2;
		list2 = list2.Next;
	}

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }

    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

    return dummyHead
}

func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
    }

	if list2 == nil {
        return list1
    }

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}
