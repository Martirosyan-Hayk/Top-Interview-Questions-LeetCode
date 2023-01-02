#include <string>
#include <stack>

class Solution
{
public:
    bool isValid(std::string s)
    {
        std::stack<char> stack;

        for (char ch : s)
        {
            if (ch == '(' || ch == '[' || ch == '{')
            {
                stack.push(ch);
            }
            else if (ch == ')' || ch == ']' || ch == '}')
            {
                if (stack.empty())
                {
                    return false;
                }
                char top = stack.top();
                if ((ch == ')' && top != '(') ||
                    (ch == ']' && top != '[') ||
                    (ch == '}' && top != '{'))
                {
                    return false;
                }
                stack.pop();
            }
        }

        return stack.empty();
    }
};