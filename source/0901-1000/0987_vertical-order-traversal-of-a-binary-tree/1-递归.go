package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode987_二叉树的垂序遍历
var m map[int][][2]int

func verticalTraversal(root *TreeNode) [][]int {
	m = make(map[int][][2]int)
	res := make([][]int, 0)
	dfs(root, 0, 0)
	arr := make([]int, 0)
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		temp := m[arr[i]]
		sort.Slice(temp, func(i, j int) bool {
			if temp[i][1] == temp[j][1] {
				return temp[i][0] < temp[j][0]
			}
			return temp[i][1] < temp[j][1]
		})
		tempArr := make([]int, 0)
		for j := 0; j < len(temp); j++ {
			tempArr = append(tempArr, temp[j][0])
		}
		res = append(res, tempArr)
	}
	return res
}

func dfs(root *TreeNode, x, y int) {
	if root == nil {
		return
	}
	m[x] = append(m[x], [2]int{root.Val, y})
	dfs(root.Left, x-1, y+1)
	dfs(root.Right, x+1, y+1)
}
