function longestCommonPrefix(strs: string[]): string {
    if (strs.length == 0) return "";
    if (strs.length == 1) return strs[0];

    let prefix: string = "";

    for (let i = 0; i < strs[0].length; i++) {
        const c: string = strs[0][i];

        for (let j = 1; j < strs.length; j++) {

            if (i >= strs[j].length || strs[j][i] != c) {
                return prefix;
            }
        }

        prefix += c;
    }

    return prefix;
};