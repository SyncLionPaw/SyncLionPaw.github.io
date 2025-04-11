package main

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。
// 如果不存在 公共子序列 ，返回 0 。

// f[i][j] 表示对(s[:i],t[:j]) 的 公共子序列的最大长度
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if text1[i-1] != text2[j-1] {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			} else {
				f[i][j] = f[i-1][j-1] + 1
			}
		}
	}
	return f[m][n]
}