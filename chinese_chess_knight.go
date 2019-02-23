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
func BFS(src *Node, dest *Node, M, N int) int {

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

		if x == dest.x && y == dest.y {
			return dist
		}

		if _, ok := visited[node]; !ok {
			visited[node] = true

			for i := 0; i < 8; i++ {
				x1 := x + row[i]
				y1 := y + col[i]

				if valid(x1, y1, M, N) {
					q.Enqueue(&Node{x1, y1, dist + 1})
				}
			}
		}
	}
	// return nil and -1 is not possible
	return -1
}

func isNodeRecord(t interface{}) bool {
	switch t.(type) {
	case *NodeRecord:
		return true
	default:
		return false
	}
}


func isNode(t interface{}) bool {
	switch t.(type) {
	case *Node:
		return true
	default:
		return false
	}
}

type NodeRecord struct {
	this *Node
	last *Node
}

func KnightFindPath(src *Node, dest *Node, M, N int) ([]*Node, int) {
	var visited = make(map[*NodeRecord]bool) // Node -> LastNode
	var q = queue.New()
	q.Enqueue(&NodeRecord{src, nil})

	for q.Len() != 0 {
		var node *NodeRecord
		if v := q.Dequeue(); isNodeRecord(v) {
			node = v.(*NodeRecord)
		}
		x := node.this.x
		y := node.this.y
		dist := node.this.dist

		if x == dest.x && y == dest.y {
			visited[node] = true
			path := findPath(src, node, visited, dist)
			return path, dist
		}

		if _, ok := visited[node]; !ok {
			visited[node] = true

			for i := 0; i < 8; i++ {
				x1 := x + row[i]
				y1 := y + col[i]

				if valid(x1, y1, M, N) {

					q.Enqueue(&NodeRecord{&Node{x1, y1, dist + 1}, node.this})
				}
			}
		}
	}
	return nil, -1
}

func findPath(src *Node, dest *NodeRecord, visited map[*NodeRecord]bool, dist int) []*Node {
	var tmpPath = make([]*Node, dist+1)
	tmpPath[0] = dest.this
	var lastNode = dest.this
	for i := 0; i < dist+1; i++ {
		if v, ok := findVisited(visited, lastNode); ok {
			tmpPath[i+1] = v
			lastNode = v
			if v.equals(*src) {
				break
			}
		} else {
			break
		}
	}
	return tmpPath
}

func findVisited(visited map[*NodeRecord]bool, lastNode *Node) (*Node, bool) {
	for k, _ := range visited {
		if k.this.x == lastNode.x && k.this.y == lastNode.y && k.this.dist == lastNode.dist {
			return k.last, true
		}
	}
	return nil, false
}
