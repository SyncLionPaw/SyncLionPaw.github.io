package main

// 装满背包的最小物品数目
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		f[i] = 10001
	}
	for _, x := range coins {
		for j := x; j <= amount; j++ {
            f[j] = min(f[j], f[j-x]+1)
		}
	}
	ans := f[amount]
	if ans >= 10001 {
		return -1
	}
	return ans
}