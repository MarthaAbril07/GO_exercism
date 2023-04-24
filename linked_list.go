package main

import (
	"fmt"
)

// Define Lista and Nodo types here.
type Nodo struct {
	Value         interface{}
	NodoAnterior  *Nodo
	NodoSiguiente *Nodo
}
type Lista struct {
	Nodos []*Nodo
}

// Note: The tests expect Nodo type to include an exported field with name Value to pass.
var nodo0 *Nodo
var nodo2 *Nodo

func NewList(args ...int) *Lista {
	lista := Lista{}
	nodos := []*Nodo{}

	for index, value := range args {

		if index == 0 {
			nodo0 = &Nodo{
				Value: args[len(args)-1],
			}
		} else {
			nodo0 = &Nodo{
				Value: args[index-1],
			}
		}

		if index == len(args)-1 {
			nodo2 = &Nodo{
				Value: args[0],
			}
		} else {
			nodo2 = &Nodo{
				Value: args[index+1],
			}
		}

		nodo1 := &Nodo{
			Value:         value,
			NodoAnterior:  nodo0,
			NodoSiguiente: nodo2,
		}
		nodos = append(nodos, nodo1)
	}
	lista.Nodos = nodos
	return &lista
}

func (n *Nodo) Next() *Nodo {
	return n.NodoSiguiente
}

func (n *Nodo) Prev() *Nodo {
	return n.NodoAnterior
}

// agregar nodo al inicio
func (l *Lista) Unshift(v int) {
	nodosActuales := l.Nodos

	nodo1 := &Nodo{
		Value:         v,
		NodoAnterior:  nodosActuales[len(nodosActuales)-1],
		NodoSiguiente: nodosActuales[0],
	}
	nodosActuales[0].NodoAnterior = nodo1
	nodosActuales[len(nodosActuales)-1].NodoSiguiente = nodo1
	nodoNuevo := []*Nodo{nodo1}
	nodoNuevo = append(nodoNuevo, nodosActuales...)
	l.Nodos = nodoNuevo
}

// agregar al final de la lista
func (l *Lista) Push(v int) {
	nodosActuales := l.Nodos
	nodo1 := &Nodo{
		Value:         v,
		NodoAnterior:  nodosActuales[len(nodosActuales)-1],
		NodoSiguiente: nodosActuales[0],
	}
	nodosActuales[0].NodoAnterior = nodo1
	nodosActuales[len(nodosActuales)-1].NodoSiguiente = nodo1
	//nodoNuevo := []*Nodo{nodo1}
	nodosActuales = append(nodosActuales, nodo1)
	l.Nodos = nodosActuales
}

// quitar el primer elemento
func (l *Lista) Shift() {
	nodosActuales := l.Nodos
	nodosActuales = append(nodosActuales[:0], nodosActuales[1:]...)
	nodosActuales[len(nodosActuales)-1].NodoSiguiente = nodosActuales[0]
	nodosActuales[0].NodoAnterior = nodosActuales[len(nodosActuales)-1]
	l.Nodos = nodosActuales
}

// eliminar el ultimo elemento
func (l *Lista) Pop() {
	/*x := []int{1, 2, 3, 4, 5}
	x = append(x[:len(x)-1], x[len(x):]...)
	fmt.Println(x)*/
	nodosActuales := l.Nodos
	nodosActuales = append(nodosActuales[:len(nodosActuales)-1], nodosActuales[len(nodosActuales):]...)
	nodosActuales[len(nodosActuales)-1].NodoSiguiente = nodosActuales[0]
	nodosActuales[0].NodoAnterior = nodosActuales[len(nodosActuales)-1]
	l.Nodos = nodosActuales
}

func (l *Lista) Reverse() {
	nuevaLista := []*Nodo{}
	for x := len(l.Nodos) - 1; x >= 0; x-- {
		item := l.Nodos[x]
		siguiente := item.NodoSiguiente
		anterior := item.NodoAnterior
		item.NodoSiguiente = anterior
		item.NodoAnterior = siguiente
		nuevaLista = append(nuevaLista, item)
	}
	l.Nodos = nuevaLista
}

func (l *Lista) First() *Nodo {
	return l.Nodos[0]
}

func (l *Lista) Last() *Nodo {
	return l.Nodos[len(l.Nodos)-1]
}

func (l *Lista) imprimirLista() {
	for index, item := range l.Nodos {
		fmt.Println("Nodo:", index)
		fmt.Println("\t[", item.NodoAnterior.Value, "<-", item.Value, "->", item.NodoSiguiente.Value)
		fmt.Println(".........")
	}
}

func main() {
	lista := NewList(1, 2, 3, 4, 5)
	//lista.imprimirLista()
	/*nodo := lista.Nodos[1]
	fmt.Println("Next:", nodo.Next().Value)
	fmt.Println("Prev:", nodo.Prev().Value)*/
	//fmt.Println("Agregando al inicio")
	//lista.Unshift(10)
	//lista.imprimirLista()
	//fmt.Println("Agregando al final")
	//lista.Push(20)
	//lista.Shift()
	//lista.Pop()
	//lista.Reverse()
	/*fmt.Println(lista.First().Value)
	fmt.Println(lista.Last().Value)*/
	lista.imprimirLista()
}
