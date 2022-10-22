function twoSum(nums, target) {
    var res = [];
    for (var i = 0; i < nums.length - 1; ++i) {
        for (var j = i + 1; j <= nums.length - 1; ++j) {
            if ((nums[i] + nums[j]) == target) {
                res[0] = i;
                res[1] = j;
                return res;
            }
        }
    }
    return res;
}
;
