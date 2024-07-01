package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 所有路径 - 从 a到b带点的所有可能路径
func allPathFromSourceToTarget(graph [][]int) (res [][]int) {
	var path []int
	var traverse func(graph [][]int, s int, path []int)

	traverse = func(graph [][]int, s int, path []int) {
		path = append(path, s)
		n := len(graph)
		if s == n-1 {
			res = append(res, append([]int(nil), path...))
			path = []int{}
		}

		// 递归相邻节点
		for _, v := range graph[s] {
			traverse(graph, v, path)
		}
	}

	traverse(graph, 0, path)

	return res
}

func main() {
	fmt.Println("1. 图从 a 点到 b 点的所有可能路径：")
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println("原图为：", graph)
	paths := allPathFromSourceToTarget(graph)
	fmt.Println("该图所有从源点到目标节点的路径为：", paths)

	line.SplitLine()

	graph1 := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println("原图为：", graph1)
	paths1 := allPathFromSourceToTarget(graph1)
	fmt.Println("该图所有从源点到目标节点的路径为：", paths1)

	line.SplitLine()

	graph2 := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println("原图为：", graph2)
	paths2 := allPathFromSourceToTarget(graph2)
	fmt.Println("该图所有从源点到目标节点的路径为：", paths2)
}
