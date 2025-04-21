package main

// 求凑成amount的方案数，完全背包
func change(amount int, coins []int) int {
	f := make([]int, amount+1) // f[i] 表示凑到i的方案数
	f[0] = 1
	for _, x := range coins {
		for j := x; j <= amount; j++ {
			f[j] += f[j-x]
		}
	}
	return f[amount]
}
