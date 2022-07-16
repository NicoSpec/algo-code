package medium

/**

给你一个字符串 s，找到 s 中最长的回文子串。

 

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
 

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
通过次数1,109,488提交次数3,009,040

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

public String longestPalindrome(String s) {
	if (s == null || s.length() < 2) {
		return s;
	}
	int strLen = s.length();
	int maxStart = 0;  //最长回文串的起点
	int maxEnd = 0;    //最长回文串的终点
	int maxLen = 1;  //最长回文串的长度

	boolean[][] dp = new boolean[strLen][strLen];

	for (int r = 1; r < strLen; r++) {
		for (int l = 0; l < r; l++) {
			if (s.charAt(l) == s.charAt(r) && (r - l <= 2 || dp[l + 1][r - 1])) {
				dp[l][r] = true;
				if (r - l + 1 > maxLen) {
					maxLen = r - l + 1;
					maxStart = l;
					maxEnd = r;

				}
			}

		}

	}
	return s.substring(maxStart, maxEnd + 1);

}