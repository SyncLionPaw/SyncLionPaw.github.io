package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
	"strings"
)

/*
给定 N 种不同面值的邮票，每种面值的邮票都有足够多张。

一个信封上最多可以贴 K 张邮票。

请你找出使得 1 到 M 之间所有的邮资都能被凑出的最大的 M 是多少。

例如，假设现在有 1 和 3 两种面值的邮票，一个信封上最多可以贴 5 张邮票，那么 1 到 13 之间的所有邮资都可以被凑出：
*/

// 使用邮票装满面值的最小数量
func stampFunc(N, K int, stamps []int) int {
	// N >= 1
	maxStamp := stamps[0]
	for _, v := range stamps {
		maxStamp = max(maxStamp, v)
	} // 找最大的面值
	maxValue := K * maxStamp     // 理论最大能凑到的面值

	f := make([]int, maxValue+1) // f[j]凑到j的最小枚数
	for i := 1; i <= maxValue; i++ {
		f[i] = math.MaxInt32 / 2
	}

	for _, x := range stamps {
		for j := 1; j <= maxValue; j++ {
			if x <= j {
				f[j] = min(f[j], f[j-x]+1) // 使用一枚，完全背包，来自同一行
			}
		}
	}
	for j := 0; j <= maxValue; j++ {
		if j == maxValue || f[j+1] > K {
			return j
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	parts := strings.Fields(scanner.Text())
	K, _ := strconv.Atoi(parts[0])
	N, _ := strconv.Atoi(parts[1])

	stamps := make([]int, 0, N)

	for len(stamps) < N {
		scanner.Scan()
		elems := strings.Fields(scanner.Text())
		for _, elem := range elems {
			a, _ := strconv.Atoi(elem)
			stamps = append(stamps, a)
		}
	}

	ans := stampFunc(N, K, stamps)
	fmt.Println(ans)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
