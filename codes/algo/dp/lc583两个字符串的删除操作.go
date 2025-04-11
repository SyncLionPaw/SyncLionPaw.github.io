// lc 583 两个字符串的删除操作
package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1) // f[i][j] 表示 word1[:i] 和 word2[:j] 相同的最小步数
	for i := range f {
		f[i] = make([]int, n+1)
	}
	// 其中一个是空串的情况
	for i := 1; i <= m; i++ {
		f[i][0] = i
	}
	for j := 1; j <= n; j++ {
		f[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				delete1 := f[i-1][j] + 1
				delete2 := f[i][j-1] + 1
				f[i][j] = min(delete1, delete2)
			}
		}
	}
	return f[m][n]
}
