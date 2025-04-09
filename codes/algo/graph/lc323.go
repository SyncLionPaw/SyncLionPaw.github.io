// 求无向图的联通分量数目
package main

import "fmt"

// 从边表 构建邻接表
func buildGraphFromEdges(n int, edges [][]int) map[int][]int {
	g := map[int][]int{}
	for i := 0; i < n; i++ {
		g[i] = []int{}
	}
	for _, tuple := range edges {
		a, b := tuple[0], tuple[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	// fmt.Printf("g: %v\n", g)
	return g
}

// 访问某个节点所在的联通分量
func graphDfs(g map[int][]int, visited *[]int, nodeId int) {
	if (*visited)[nodeId] > 0 { // 已经访问过
		return
	}
	(*visited)[nodeId]++
	for _, conn := range g[nodeId] {
		graphDfs(g, visited, conn)
	}
}

func countComponents(n int, edges [][]int) int {
	g := buildGraphFromEdges(n, edges)
	visited := make([]int, n) // 标记每个节点是否被访问过

	componentCnt := 0

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			componentCnt++
			graphDfs(g, &visited, i)
		}
	}
	return componentCnt
}

func countComponents2(n int, edges [][]int) int {
	adjcent := make([][]int, n) // 领接表不一定需要使用哈希表，直接使用二维数组也可以
	for i := 0; i < n; i++ {
		adjcent[i] = []int{}
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjcent[a] = append(adjcent[a], b)
		adjcent[b] = append(adjcent[b], a)
	}
	componentCnt := 0

	var dfs func(adjcent [][]int, id int, visited *[]int)
	dfs = func(adjcent [][]int, id int, visited *[]int) {
		if (*visited)[id] == 1 {
			return
		}
		(*visited)[id] = 1
		for _, conn := range adjcent[id] {
			dfs(adjcent, conn, visited)
		}
	}

	visited := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			componentCnt++
			dfs(adjcent, i, &visited)
		}
	}
	return componentCnt
}
func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	cnt := countComponents(n, edges)
	fmt.Printf("cnt: %v\n", cnt)

	n = 5
	edges = [][]int{{0, 1}, {1, 2}, {3, 4}}
	cnt = countComponents(n, edges)
	fmt.Printf("cnt: %v\n", cnt)
}
