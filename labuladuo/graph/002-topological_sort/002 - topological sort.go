package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 课程表 - 环检测
// 你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程 bi 。
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false
var (
	hasCycle bool
	visited  map[int]bool
	onPath   map[int]bool
)

func courseCanFinished(numCourses int, prerequisites [][]int) bool {
	visited = make(map[int]bool)
	onPath = make(map[int]bool)
	graph := buildGraph(numCourses, prerequisites)
	for k, _ := range graph {
		traverseCourseCanFinished(graph, k)
	}

	return !hasCycle
}

func traverseCourseCanFinished(graph map[int][]int, s int) {
	if onPath[s] {
		hasCycle = true
	}

	// 不需要继续遍历了
	if res := visited[s]; res || onPath[s] {
		return
	}

	visited[s] = true
	onPath[s] = true
	for _, v := range graph[s] {
		traverseCourseCanFinished(graph, v)
	}
	onPath[s] = false
}

// 构建 邻接表
func buildGraph(numsCourses int, prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)

	// init
	for i := 0; i < numsCourses; i++ {
		graph[i] = []int{}
	}

	for _, v := range prerequisites {
		from, to := v[0], v[1]
		graph[from] = append(graph[from], to)
	}

	return graph
}

// 课程表2 - 拓扑排序算法
// 总共有 n 门课，记为 0 - n-1, [0, 1] 表示选修 0 需要先完成 1
// 返回为完成课程学习，所安排的学习顺序
func courseFindOrder(numsCourse int, prerequisites [][]int) []int {
	if !courseCanFinished(numsCourse, prerequisites) {
		return []int{}
	}

	return []int{}
}

func main() {
	fmt.Println("1. 课程表 - 是否能完成所有可能的学习(DFS)：")
	numsCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	finished := courseCanFinished(numsCourses, prerequisites)
	fmt.Println("判责的结果为：", finished)

	line.SplitLine()
}
