function isValid(s) {
    let stack = [];
    for (let ch of s) {
        if (ch == '(' || ch == '[' || ch == '{') {
            stack.push(ch);
        }
        else if (ch == ')' || ch == ']' || ch == '}') {
            if (stack.length == 0) {
                return false;
            }
            let top = stack.pop();
            if ((ch == ')' && top != '(') ||
                (ch == ']' && top != '[') ||
                (ch == '}' && top != '{')) {
                return false;
            }
        }
    }
    return stack.length == 0;
}