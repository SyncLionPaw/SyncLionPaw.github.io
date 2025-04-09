// 算法设计与分析第五版 算法实现题4-1
package main

import (
	"fmt"
	"sort"
)

// 求安排活动所需要的最少会场数
func greedyActivity(activities [][2]int) int {
	place := 0
	n := len(activities)
	sort.Slice(activities, func(i, j int) bool {
		if activities[i][0] < activities[j][0] {
			return true
		} else if activities[i][0] > activities[j][0] {
			return false
		} else {
			return activities[i][1] < activities[j][1]
		}
	})
}