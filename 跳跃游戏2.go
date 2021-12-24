package main

/**
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

 

示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 

提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
// 思路 动态规划法 dp[i] 表示 i下标能不能到达
// dp[i] 能不能到达 取决于 dp[0:i-1] 中能到达的地方能不能到达dp[i]
//  dp[j] 为 dp[i]  前的某个位置  那么当 dp[j]=true 并且 nums[j]>=i-j 那么 i就能到达
func canJump(nums []int) bool {
	if len(nums)<=1{
		return true
	}
	dp:=make([]bool,len(nums))
	// 第一个肯定能到达(起始条件)
	dp[0]=true
	for i:=1;i<len(nums);i++{
		for j:=i-1;j>=0;j--{
			if dp[j] && nums[j]>=i-j{
				dp[i]=true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
