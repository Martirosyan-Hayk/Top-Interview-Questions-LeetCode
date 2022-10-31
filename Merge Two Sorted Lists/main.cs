

public class ListNode
{
    public int val;
    public ListNode next;
    public ListNode(int val = 0, ListNode next = null)
    {
        this.val = val;
        this.next = next;
    }
}

public class Solution
{
    public ListNode MergeTwoLists(ListNode list1, ListNode list2)
    {
        if (list1 == null || list2 == null)
            return list1 == null ? list2 : list1;

        ListNode head;
        ListNode curr;

        if (list1.val < list2.val)
        {
            curr = list1;
            head = list1;
            list1 = list1.next;
        }
        else
        {
            curr = list2;
            head = list2;
            list2 = list2.next;
        }

        while (list1 != null && list2 != null)
        {
            if (list1.val < list2.val)
            {
                curr.next = list1;
                list1 = list1.next;
            }
            else
            {
                curr.next = list2;
                list2 = list2.next;
            }
            curr = curr.next;
        }

        if (list1 != null)
            curr.next = list1;
        else
            curr.next = list2;

        return head;
    }

    public ListNode MergeTwoListsRecursive(ListNode list1, ListNode list2)
    {
        if (list1 == null)
            return list2;

        if (list2 == null)
            return list1;

        if (list1.val <= list2.val)
        {
            list1.next = MergeTwoListsRecursive(list1.next, list2);
            return list1;
        }
        else
        {
            list2.next = MergeTwoListsRecursive(list1, list2.next);
            return list2;
        }
    }
}
