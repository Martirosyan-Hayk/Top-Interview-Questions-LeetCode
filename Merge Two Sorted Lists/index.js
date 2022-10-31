/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function (list1, list2) {
    if (!list1 || !list2)
        return !list1 ? list2 : list1;

    let head;
    let curr;

    if (list1.val < list2.val) {
        curr = list1;
        head = list1;
        list1 = list1.next;
    }
    else {
        curr = list2;
        head = list2;
        list2 = list2.next;
    }

    while (list1 && list2) {
        if (list1.val < list2.val) {
            curr.next = list1;
            list1 = list1.next;
        }
        else {
            curr.next = list2;
            list2 = list2.next;
        }
        curr = curr.next;
    }

    if (list1)
        curr.next = list1;
    else
        curr.next = list2;

    return head;
};

var mergeTwoListsRecursive = function (list1, list2) {
    if (!list1)
        return list2;

    if (!list2)
        return list1;

    if (list1.val <= list2.val) {
        list1.next = mergeTwoListsRecursive(list1.next, list2);
        return list1;
    }
    else {
        list2.next = mergeTwoListsRecursive(list1, list2.next);
        return list2;
    }
};