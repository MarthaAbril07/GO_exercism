package main

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

type OrdenarPorParent []Record

func (a OrdenarPorParent) Len() int {
	return len(a)
}

func (a OrdenarPorParent) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a OrdenarPorParent) Less(i, j int) bool {
	return a[i].Parent < a[j].Parent
}

type OrdenarPorID []Record

func (a OrdenarPorID) Len() int {
	return len(a)
}

func (a OrdenarPorID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a OrdenarPorID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

// (*Nodo, error)
/*func Build(records []Record) (*Nodo, error) {
	nodoPrincipal := Nodo{}
	var err error
	if len(records) > 0 {
		sort.Sort(OrdenarPorParent(records))
		var cont = 0
		for len(records) > 0 {
			item := records[cont]
			padre := item.Parent
			nodo := &Nodo{
				ID: item.ID,
			}
			if item.ID == 0 && item.Parent == 0 {
				nodoPrincipal.ID = 0
				nodo.Children = []*Nodo{}
			} else {
				nodoPrincipal = *agregar(padre, nodo, &nodoPrincipal)
			}
			records = append(records[:cont], records[cont+1:]...)
		}
		imprimir(&nodoPrincipal)
		return &nodoPrincipal, err
	} else {
		err = errors.New("Error")
		return nil, err
	}
}*/

func Build(records []Record) (*Node, error) {
	nodoPrincipal := Node{}
	var err error = nil

	var cont = 0
	if len(records) > 0 {
		if len(records) == 1 {
			nodoPrincipal.ID = records[0].Parent
		} else {
			sort.Sort(OrdenarPorParent(records))
			max := records[len(records)-1].Parent
			//nodoPrincipal.ID = 0
			for cont <= max {
				nodos := getNodos(records, cont)
				nodoPrincipal = *agregar(cont, nodos, &nodoPrincipal)
				cont++
			}
		}

		return &nodoPrincipal, err
	} else {
		err = errors.New("Error")
		return nil, err
	}
}

func getNodos(records []Record, cont int) []*Node {
	nodos := []*Node{}
	sort.Sort(OrdenarPorID(records))
	for _, item := range records {
		nodo := &Node{
			ID: item.ID,
		}
		if item.Parent == 0 && item.ID == 0 {

		} else {
			if item.Parent == cont {
				nodos = append(nodos, nodo)
			}
		}
	}
	return nodos
}

func agregar(padre int, hijos []*Node, nodoPrincipal *Node) *Node {
	if nodoPrincipal.ID == padre {
		for _, item := range hijos {
			nodoPrincipal.Children = append(nodoPrincipal.Children, item)
		}
	} else {
		for _, item := range nodoPrincipal.Children {
			agregar(padre, hijos, item)
		}
	}
	return nodoPrincipal
}

func imprimir(nodo *Node) {
	if len(nodo.Children) > 0 {
		fmt.Println("ID-padre", nodo.ID)
		for _, y := range nodo.Children {
			imprimir(y)
		}
	} else {
		fmt.Println("\t\tID-hijo:", nodo.ID)
	}
}

func main() {
	records := []Record{
		{ID: 0, Parent: 1},
	}
	nodos, _ := Build(records)
	imprimir(nodos)
}
