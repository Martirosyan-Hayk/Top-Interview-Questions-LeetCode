#include <vector>

class Solution
{
public:
    std::vector<int> twoSum(std::vector<int> &nums, int target)
    {
        std::vector<int> res;

        int size = nums.size() - 1;

        for (int i = 0; i < size; ++i)
        {
            for (int j = i + 1; j <= size; ++j)
            {
                if ((nums[i] + nums[j]) == target)
                {
                    res.push_back(i);
                    res.push_back(j);
                    return res;
                }
            }
        }

        return res;
    }
};
