using System.Collections.Generic;


public class Solution
{
    public int RomanToInt(string s)
    {
        Dictionary<char, int> romanNumbers =
                      new Dictionary<char, int>();
        romanNumbers['I'] = 1;
        romanNumbers['V'] = 5;
        romanNumbers['X'] = 10;
        romanNumbers['L'] = 50;
        romanNumbers['C'] = 100;
        romanNumbers['D'] = 500;
        romanNumbers['M'] = 1000;

        int number = 0;

        for (int i = 0; i < s.Length; ++i)
        {
            if (i + 1 < s.Length && romanNumbers[s[i]] < romanNumbers[s[i + 1]])
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
}