package main

import (
	"AttackGraph/topo"
	"fmt"
)

func main() {
	var topo topo.NetTopo

	err := topo.SetAdjacentMatrix([][]bool{
		{true, true, true, false, true},
		{false, true, false, true, false},
		{false, false, true, true, false},
		{false, false, false, true, false},
		{false, true, false, false, true},
	}, []string{"h1", "h2", "h3", "h4", "h5"})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("Adjacent Matrix:")
	topo.PrintAdjacentMatrix()

}
