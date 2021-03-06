package main

import "fmt"

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

 

示例 1：


输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6
 

提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
func main() {
	num := uniquePath(7,3)
	fmt.Println(num)
}
func uniquePath(m, n int) int {
	num := uniquePathsHelper(1, 1, m, n)
	return num
}


//第i行第j列到第m行第n列共有多少种路径
func uniquePathsHelper(i,  j, m,  n int)int {
	//边界条件的判断
	if i > m || j > n {
		return 0

	}
	if i == m && j == n {
		return 1
	}
	//从右边走有多少条路径
	right := uniquePathsHelper(i + 1, j, m, n)
	//从下边走有多少条路径
	down := uniquePathsHelper(i, j + 1, m, n)
	//返回总的路径
	return right + down
}
