var longestCommonPrefix = function (strs) {
    if (strs.length == 0) return "";
    if (strs.length == 1) return strs[0];

    let prefix = "";

    for (let i = 0; i < strs[0].length; i++) {
        const c = strs[0][i];

        for (let j = 1; j < strs.length; j++) {

            if (i >= strs[j].length || strs[j][i] != c) {
                return prefix;
            }
        }

        prefix += c;
    }

    return prefix;
};