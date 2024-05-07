package node

const TYPE_VULN = 0
const TYPE_PRIILEGE = 1

/*
This is a attribute node struct, the node has two type: the vulnability or the privilege.
The Type is mark what type the node is, the Name is the display name in the graph.
*/
type AttributeNode struct {
	Type int
	Name string
}

func (ac AttributeNode) Hash() string {
	return ac.Name
}

func (an *AttributeNode) SetName(name string) {
	an.Name = name
}
