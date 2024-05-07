package attackgraph

import (
	"AttackGraph/node"
	"AttackGraph/topo"
	vulnagent "AttackGraph/vulnAgent"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type AttackGraph[N node.Node] struct {
	Name string
	g    graph.Graph[string, N]
}

func New[N node.Node](name string, options ...func(*graph.Traits)) AttackGraph[N] {
	t := AttackGraph[N]{Name: name}
	t.g = graph.New(N.Hash, options...)
	return t
}

func (graph *AttackGraph[N]) AddVertex(n N, options ...func(*graph.VertexProperties)) error {
	return graph.g.AddVertex(n, options...)
}

func (graph *AttackGraph[N]) AddEdge(s N, t N, options ...func(*graph.EdgeProperties)) error {
	return graph.g.AddEdge(s.Hash(), t.Hash(), options...)
}

func (graph *AttackGraph[N]) GenerateAttackGraph(topo topo.NetTopo, vulns []vulnagent.CVE) error {
	panic("Not implement function: GenerateAttackGraph")
}

func (graph AttackGraph[N]) Show() error {
	fileName := graph.Name + ".gv"
	file, _ := os.Create(fileName)
	err := draw.DOT(graph.g, file)
	if err != nil {
		return err
	}

	cmd := exec.Command("dot", "-Tpng", "-O", fileName)
	_, err = cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
