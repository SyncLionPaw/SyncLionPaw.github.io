// lc 72 编辑距离问题
// https://www.hello-algo.com/chapter_dynamic_programming/edit_distance_problem/
package main

func minDistance(s, t string) int {
	m, n := len(s), len(t)
	f := make([][]int, m+1) // f[i][j] 表示s[i]变成t[j] 的最少编辑次数
	for i := range f {
		f[i] = make([]int, n+1)
	}
	// 初始化边界, 其中一个字符串为空的情况
	for i := 1; i <= m; i++ {
		f[i][0] = i
	}
	for j := 1; j <= n; j++ {
		f[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i] == t[j] {
				f[i][j] = f[i-1][j-1]
			} else { // s[i] != t[j]
				append := f[i][j-1] + 1   // 在s[i-1]后面直接加上t[j-1]
				change := f[i-1][j-1] + 1 // 把s[i-1]替换为t[j-1]
				delete := f[i-1][j] + 1   // 删除s[i-1]
				f[i][j] = min(append, change, delete)
			}
		}
	}
	return f[m][n]
}
