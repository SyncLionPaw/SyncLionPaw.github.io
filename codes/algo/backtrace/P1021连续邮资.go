package main

import (
	"fmt"
	"math"
)

func maxCovers(stamps []int, N int) (r []int) {
	stampType := len(stamps)
	maxValue := stamps[stampType-1] * N  // 理论能达到的最大面值
	minUse := make([][]int, stampType)

	for i := range minUse {
		minUse[i] = make([]int, maxValue+1)
		for j := range minUse[i] {
			minUse[i][j] = math.MaxInt64
		}
	}

	for j := 0; j < maxValue; j++ {
		if j*1 > N {
			break
		}
		minUse[0][j] = j
	}

	for i := 0; i < stampType; i++ {
		minUse[i][0] = 0
	}

	for i := 0; i < stampType; i++ {
		for j := 1; j < maxValue; j++ {
			curMin := minUse[i][j]
			for t := 0; t < N; t++ {
				if j < t*stamps[i] {
					break
				}
				if i > 0 {
					if minUse[i-1][j-t*stamps[i]] != math.MaxInt64 {
						if t+minUse[i-1][j-t*stamps[i]] < curMin {
							curMin = t + minUse[i-1][j-t*stamps[i]]
						}
					}
				}
			}
			minUse[i][j] = curMin
		}
	}

	r = make([]int, stampType)

	for i := 0; i < stampType; i++ {
		for j := 0; j < maxValue; j++ {
			if j+1 < len(minUse[i]) && minUse[i][j+1] > N && minUse[i][j] <= N {
				r[i] = j
			}
		}
	}
	return
}

var (
	ans   int
	stamps []int
)

func backtrace(i int, path []int) {
	tmp := maxCovers(path, N)
	globalR := tmp[len(tmp)-1]
	if i == K {
		if globalR >= ans {
			ans = globalR
			stamps = make([]int, len(path))
			copy(stamps, path)
		}
		return
	}

	for v := path[len(path)-1] + 1; v <= globalR+1; v++ {
		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, v)
		backtrace(i+1, newPath)
	}
}

var N, K int

func main() {
	fmt.Scan(&N, &K)
	path := []int{1}
	ans = 1
	stamps = []int{}

	backtrace(1, path)
	for i, v := range stamps {
		fmt.Printf("%v", v)
		if i == len(stamps)-1 {
			fmt.Println()
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("MAX=%d\n", ans)
}