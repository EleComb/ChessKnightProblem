package ChessKnightProblem

type Node struct {
	x    int
	y    int
	dist int
}

func (n *Node) setNode(x, y int) {
	n.x = x
	n.y = y
}

func (n *Node) setNodeDist(x, y, dist int) {
	n.x = x
	n.y = y
	n.dist = dist
}

func (n *Node) equals(i Node) bool {
	if *n == i {
		return true
	}
	if n.x != i.x {
		return false
	}
	if n.y != i.y {
		return false
	}
	return n.dist == i.dist
}

func (n *Node) hashCode() int {
	var result = n.x
	result = 31 * result + n.y
	result = 31 * result + n.dist
	return result
}


