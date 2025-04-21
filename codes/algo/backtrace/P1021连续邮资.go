// https://www.luogu.com.cn/problem/P1021
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func stampFunc(K int, stamps []int) int {
	maxStamp := stamps[0]
	for _, v := range stamps {
		maxStamp = max(maxStamp, v)
	}
	maxValue := K * maxStamp

	f := make([]int, maxValue+2)
	for i := 1; i <= maxValue; i++ {
		f[i] = math.MaxInt32 / 2
	}

	for _, x := range stamps {
		for j := x; j <= maxValue; j++ {
			if f[j-x]+1 < f[j] {
				f[j] = f[j-x] + 1
			}
		}
	}

	for j := 1; j <= maxValue; j++ {
		if f[j] > K {
			return j - 1
		}
	}
	return maxValue
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())

	N, _ := strconv.Atoi(parts[0])
	K, _ := strconv.Atoi(parts[1]) // 修正K的获取

	path := []int{1}
	ans_stamps := []int{}
	ans := 1

	var backtrace func(i int)
	backtrace = func(i int) {
		r := stampFunc(N, path)
		// fmt.Printf("r: %v\n", r)
		// fmt.Printf("path: %v\n", path)
		if i == K { // K 种邮票，修正终止条件
			if ans < r {
				ans = r
				tmp := make([]int, len(path))
				copy(tmp, path)
				ans_stamps = tmp
			}
			return
		}
		for j := path[i-1] + 1; j <= r+1; j++ {
			path = append(path, j)
			backtrace(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrace(1)
	for i, x := range ans_stamps {
		fmt.Printf("%v", x)
		if i == len(ans_stamps)-1 {
			fmt.Println()
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("MAX=%d\n", ans)
}
