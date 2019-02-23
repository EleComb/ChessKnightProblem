package ChessKnightProblem

import (
	"github.com/golang-collections/collections/queue"
)

var row = []int{2, 2, -2, -2, 1, 1, -1, -1}
var col = []int{-1, 1, 1, -1, 2, -2, 2, -2}

// in chinese chess
// the size of the chessboard is 10*9
func valid(x, y, M, N int) bool {
	//if x < 0 || y < 0 || x >= N || y >= N {
	//	return false
	//}
	//return true
	return !(x < 0 || y < 0 || x >= M || y >= N)
}

// Find minimum number of steps taken by the knight
// from source to reach destination using BFS
func BFS(src *Node, dest *Node, M, N int) ([]*Node, int) {
	var nodeQuque = make([]*Node, 0)

	var visited = make(map[*Node]bool)
	var q = queue.New()
	q.Enqueue(src)

	for q.Len() != 0 {
		var node *Node
		if v := q.Dequeue(); isNode(v) {
			node = v.(*Node)
		}
		x := node.x
		y := node.y
		dist := node.dist
		nodeQuque = append(nodeQuque, node)

		if x == dest.x && y == dest.y {
			return nil, dist
		}

		if _, ok := visited[node]; !ok {
			visited[node] = true

			for i := 0; i < 8; i++ {
				x1 := x + row[i]
				y1 := y + col[i]

				if valid(x1, y1, M, N) {
					q.Enqueue(&Node{x1, y1, dist+1})
				}
			}
		}
	}
	// return nil and -1 is not possible
	return nil, -1
}

func isNode(t interface{}) bool {
	switch t.(type) {
	case *Node:
		return true
	default:
		return false
	}
}