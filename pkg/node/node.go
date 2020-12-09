package node

import (
	"fmt"
	"reflect"
)

type Node struct {
	Typ      NodeType
	In       []interface{}
	Inserter bool
}

func NewNode(nodeType NodeType, in []interface{}) Node {
	for i, r := range in {
		if reflect.TypeOf(r) != nodeType.In[i] {
			panic(fmt.Sprintf("The type is wrong on input %d ", i))
		}
	}
	return Node{
		Typ: nodeType,
		In:  in,
	}
}

func (n Node) Run() []interface{} {
	return n.Typ.Run(n.In)
}

func FastRun(nodeType NodeType, in []interface{}) []interface{} {
	return NewNode(nodeType, in).Run()
}
