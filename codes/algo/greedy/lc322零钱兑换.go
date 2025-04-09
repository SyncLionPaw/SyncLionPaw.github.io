// 322 [零钱兑换](https://leetcode.cn/problems/coin-change/description/)

package main

import (
	"fmt"
	"math"
)

// 动态规划法
func coinChange(coins []int, amount int) int {
	n := len(coins)
	// f[i][j] 表示使用前i种硬币，凑齐j元所用的最少数量，这里是完全背包
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2
		}
	}
	for i := range f {
		f[i][0] = 0
	}
	for i, x := range coins {
		for j := 1; j < amount+1; j++ {
			if x > j {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = min(f[i][j], f[i+1][j-x]+1)
			}
		}
	}
	ans := f[n][amount]
	if ans >= math.MaxInt/2 {
		return -1
	}
	return ans
}

// 对于某些硬币面值组合，贪心算法并不能找到最优解
// 零钱兑换问题，贪心算法无法保证找到全局最优解，并且有可能找到非常差的解。它更适合用动态规划解决。
func coinChangeGreedy(coins []int, amount int) int {
	//贪心地选择不大于且最接近它的硬币
	coinCnt := 0
	index := len(coins) - 1 // 从最后一种硬币开始枚举
	for index >= 0 && amount > 0 {
		if coins[index] <= amount {
			amount -= coins[index]
			coinCnt++
		} else {
			index--
		}
	}
	if amount == 0 {
		return coinCnt
	}
	return -1
}

func main() {
	coins := []int{1, 5, 10, 20, 50}
	ans := coinChangeGreedy(coins, 131)
	fmt.Printf("ans: %v\n", ans) // 5，对于这种组合 能得到最优解

	coins = []int{1, 20, 50}
	ans = coinChangeGreedy(coins, 60)
	fmt.Printf("对于%v硬币组合, 无法得到最优解: %v\n", coins, ans) // 这种币值组合，无法得到最优解

	correct := coinChange(coins, 60)
	fmt.Printf("dp能求最优解: %v\n", correct) // dp能求最优解
}
