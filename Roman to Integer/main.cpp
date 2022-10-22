#include <string>
#include <map>

class Solution
{
public:
    int romanToInt(std::string s)
    {
        std::map<char, int> romanNumbers;
        romanNumbers['I'] = 1;
        romanNumbers['V'] = 5;
        romanNumbers['X'] = 10;
        romanNumbers['L'] = 50;
        romanNumbers['C'] = 100;
        romanNumbers['D'] = 500;
        romanNumbers['M'] = 1000;

        int number = 0;

        for (int i = 0; i < s.size(); ++i)
        {
            if (romanNumbers[s[i]] < romanNumbers[s[i + 1]])
            {
                number -= romanNumbers[s[i]];
            }
            else
            {
                number += romanNumbers[s[i]];
            }
        }
        return number;
    }
};