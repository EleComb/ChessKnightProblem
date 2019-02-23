package ChessKnightProblem

import (
	"fmt"
	"testing"
)

// https://www.techiedelight.com/chess-knight-problem-find-shortest-path-source-destination/
// https://www.geeksforgeeks.org/minimum-steps-reach-target-knight/
// https://www.geeksforgeeks.org/minimum-steps-reach-target-knight-set-2/
// this time I copy THE JAVA VERSION
func TestChineseChessKnight(t *testing.T) {
	for _, unit := range []struct {
		source      *Node
		destination *Node
	}{
		//{
		//	&Node{0, 7, 0},
		//	&Node{7, 0, 0},
		//},
		{
			&Node{4, 5, 0},
			&Node{1, 1, 0},
		},
	}{
		path, length := BFS(unit.source, unit.destination, 10, 9)
		fmt.Println(path, length)
	}
}
