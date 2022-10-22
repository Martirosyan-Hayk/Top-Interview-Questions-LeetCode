function romanToInt(s) {
    var romanNumbers = new Map();
    romanNumbers['I'] = 1;
    romanNumbers['V'] = 5;
    romanNumbers['X'] = 10;
    romanNumbers['L'] = 50;
    romanNumbers['C'] = 100;
    romanNumbers['D'] = 500;
    romanNumbers['M'] = 1000;
    var number = 0;
    for (var i = 0; i < s.length; ++i) {
        if (romanNumbers[s[i]] < romanNumbers[s[i + 1]]) {
            number -= romanNumbers[s[i]];
        }
        else {
            number += romanNumbers[s[i]];
        }
    }
    return number;
}
;
function romanToInt2(s) {
    var romanNumbers = new Map();
    romanNumbers.set('I', 1);
    romanNumbers.set('V', 5);
    romanNumbers.set('X', 10);
    romanNumbers.set('L', 50);
    romanNumbers.set('C', 100);
    romanNumbers.set('D', 500);
    romanNumbers.set('M', 1000);
    var number = 0;
    for (var i = 0; i < s.length; ++i) {
        var char1 = romanNumbers.get(s[i]);
        var char2 = romanNumbers.get(s[i + 1]);
        if (char1 && char2 && char1 < char2) {
            number -= char1;
        }
        else if (char1) {
            number += char1;
        }
    }
    return number;
}
;
