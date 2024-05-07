package topo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type NetTopo struct {
	Name        string
	tables      [][]string //邻接表
	matrix      [][]bool   //邻接矩阵
	allNodeName []string   //设备名称
}

type NetTopoError struct {
	msg string
}

func (nte *NetTopoError) Error() string {
	return nte.msg
}

func (t *NetTopo) AddNode(nodeName string) {
	t.allNodeName = append(t.allNodeName, nodeName)
	t.tables = append(t.tables, []string{nodeName})
}

func (t *NetTopo) AddNodeAdjacentNode(host string, adjacentNode []string) error {
	index, err := t.getIndex(host)
	if err != nil {
		return err
	}
	if t.tableBoundCheck(index) != nil {
		return err
	}
	t.tables[index] = append(t.tables[index], adjacentNode...)
	return nil
}

func (t NetTopo) PrintAdjacentTable() {
	if len(t.tables) == 0 {
		fmt.Println("Table has no content")
	}
	for _, adjacentNode := range t.tables {
		fmt.Println(strings.Join(adjacentNode, " -> "))
	}
}

func (t NetTopo) getIndex(nodeName string) (int, error) {
	for index, value := range t.allNodeName {
		if value == nodeName {
			return index, nil
		}
	}
	return -1, &NetTopoError{"Cann't find the node in All Node Name"}
}

func (t NetTopo) tableBoundCheck(index int) error {
	if len(t.tables) < index {
		return &NetTopoError{"Out of range"}
	}
	return nil
}

func (t NetTopo) GetAdjacentNodeByName(nodeName string) ([]string, error) {

	index, err := t.getIndex(nodeName)
	if err != nil {
		return nil, err
	}

	err = t.tableBoundCheck(index)
	if err != nil {
		return nil, err
	}

	return t.tables[index], nil
}

func (t NetTopo) GetAdjacentNodeByIndex(index int) ([]string, error) {

	err := t.tableBoundCheck(index)
	if err != nil {
		return nil, err
	}

	return t.tables[index], nil
}

func (t NetTopo) IsInTopo(nodeName string) bool {

	for _, value := range t.allNodeName {
		if value == nodeName {
			return true
		}
	}
	return false
}

/*
To set the topo as a adjacent matrix, provide the complete Matrix and the Nodes name, the node's name in nodeName is assign to the Matrix colum.
ATTENTION: use this function will override the data that topo already setted, and it will clear the adjacent table. Please use this function just in init.
*/
func (t *NetTopo) SetAdjacentMatrix(matrix [][]bool, nodeName []string) error {
	// check the matrix is a square
	if len(matrix) != len(nodeName) {
		return &NetTopoError{"The Matrix scale doesn't match the node numbers"}
	}
	for index, value := range matrix {
		if len(value) != len(nodeName) {
			return &NetTopoError{fmt.Sprintf("The colum %d length doesn't match the node numbers", index)}
		}
	}

	//All check passed
	t.matrix = matrix
	t.allNodeName = nodeName
	t.tables = [][]string{}

	return nil
}

/*
Generate the adjacent table from the adjacent matrix.
*/
func (t *NetTopo) GenerateAdjacentTable(override bool) error {
	if !override && len(t.tables) != 0 {
		return &NetTopoError{"The table is all ready setted"}
	}

	for host_index, host_name := range t.allNodeName {
		tmp := []string{host_name}
		for index, value := range t.matrix[host_index] {
			if host_index == index {
				continue
			}
			if value {
				tmp = append(tmp, t.allNodeName[index])
			}
		}
		t.tables = append(t.tables, tmp)
	}

	return nil
}

func (t NetTopo) PrintAdjacentMatrix() {
	if len(t.matrix) == 0 {
		fmt.Println("Matrix has no content")
		return
	}
	if len(t.allNodeName) == 0 {
		fmt.Println("No node in topo")
		return
	}

	max := func(s []string) int {
		m := -1
		for _, v := range s {
			if len(v) > m {
				m = len(v)
			}
		}
		return m
	}

	space := max(t.allNodeName) + 4 // space is the unit length for print
	align := func(s string) string {
		fs := "%" + strconv.Itoa(space) + "s"
		return fmt.Sprintf(fs, s)
	}

	//build head
	tmp := []string{align(" ")}
	for _, v := range t.allNodeName {
		tmp = append(tmp, align(v))
	}
	fmt.Println(strings.Join(tmp, ""))

	// body
	b2i := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	for i, v := range t.allNodeName {
		tmp = []string{align(v)}
		for _, connected := range t.matrix[i] {
			tmp = append(tmp, align(strconv.Itoa(b2i(connected))))
		}
		fmt.Println(strings.Join(tmp, ""))
	}

}

func (t *NetTopo) Show() error {
	g := graph.New(graph.StringHash, graph.Directed())
	for _, name := range t.allNodeName {
		err := g.AddVertex(name)
		if err != nil {
			return err
		}
	}

	// priority use adjacent table
	if len(t.tables) != 0 {
		for i, host := range t.allNodeName {
			for _, name := range t.tables[i][1:] {
				err := g.AddEdge(host, name)
				if err != nil {
					return err
				}
			}
		}
	} else {
		for i, host := range t.allNodeName {
			for node, v := range t.matrix[i] {
				if i == node {
					continue
				}
				if v {
					err := g.AddEdge(host, t.allNodeName[node])
					if err != nil {
						return err
					}
				}
			}
		}
	}
	fileName := t.Name + ".gv"
	file, _ := os.Create(fileName)
	err := draw.DOT(g, file)
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
