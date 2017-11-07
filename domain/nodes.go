package domain

type Node struct {
	LOC      int //LinesOfCode
	Path     string
	Children []*Node
	Parents  []Node
}

type Project Node

func (p *Project) AddLoc(loc int) {
	p.LOC += loc
}

func (n *Node) AddChild(node Node) {
	n.Children = append(n.Children, &node)
}

func (n *Node) AddParent(node Node) {
	n.Parents = append(n.Parents, node)
}

func (n *Node) getAllParentsLoc() int {
	loc := 0
	toProcess := n.Parents

	for i := 0; i < len(toProcess); i++ {
		for _, grandParent := range toProcess[i].Parents {
			if !grandParent.isIn(toProcess) {
				toProcess = append(toProcess, grandParent)
			}
		}
		loc += toProcess[i].LOC
	}

	return loc
}
func (n *Node) getAllParentsCount() int {
	toProcess := n.Parents

	for i := 0; i < len(toProcess); i++ {
		for _, grandParent := range toProcess[i].Parents {
			if !grandParent.isIn(toProcess) {
				toProcess = append(toProcess, grandParent)
			}
		}
	}
	return len(toProcess)
}

func (n Node) isIn(arr []Node) bool {
	for _, arrN := range arr {
		if arrN.Path == n.Path {
			return true
		}
	}
	return false
}

func (n *Node) getChildrenLoc() int {
	loc := 0
	for _, child := range n.Children {
		loc += child.LOC
	}
	return loc
}
