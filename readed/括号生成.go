package main

import "fmt"

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
func main() {
	res := generateParenthesis(5)
	fmt.Println(res)
}
func generateParenthesis(n int) []string {
	res := []string{}

	digui(1,n,0,1,"",&res)
	return res
}

// index 第几个括号, 用于退出条件，当index == 2*n 要退出
// n 给的n
// sum ,如果输入 ( 是+1，）是-1 那当全部括号输入完毕，sum必须==0
// add ,输入左括号或者右括号, +1或者-1
// str 每次添加括号后的字符串
// res 结果切片，这里使用地址，防止结果丢失
func digui(index,n,sum,add int, str string ,res *[]string) {
	sum = sum + add
	//不能出现sum<0的情况
	if sum < 0 {
		return
	}

	if add == 1 {
		str += "("
	} else {
		str += ")"
	}

	if index == n*2 {

		if sum == 0{
			*res = append(*res,str)
		}
		return
	}

	tempIndex := index+1
	//添加下一个括号有两种可能，
	//第一种是添加 (
	//第二种是添加 )
	digui(tempIndex,n,sum,1,str,res)
	digui(tempIndex,n,sum,-1,str,res)
}