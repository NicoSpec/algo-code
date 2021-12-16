package main

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
 

提示：

1 <= n <= 8


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

class Solution {
public List<String> generateParenthesis(int n) {
List<String> combinations = new ArrayList<String>();
generateAll(new char[2 * n], 0, combinations);
return combinations;
}

public void generateAll(char[] current, int pos, List<String> result) {
if (pos == current.length) {
if (valid(current)) {
result.add(new String(current));
}
} else {
current[pos] = '(';
generateAll(current, pos + 1, result);
current[pos] = ')';
generateAll(current, pos + 1, result);
}
}

public boolean valid(char[] current) {
int balance = 0;
for (char c: current) {
if (c == '(') {
++balance;
} else {
--balance;
}
if (balance < 0) {
return false;
}
}
return balance == 0;
}
}