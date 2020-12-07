package node

import (
	"fmt"
)

type Node struct {
	Typ NodeType
	in  []interface{}
}

func NewNode(nodeType NodeType, in []interface{}) Node {
	for i, r := range in {
		if Rtype(r) != nodeType.In[i] {
			panic(fmt.Sprintf("The type is wrong on input %d ", i))
		}
	}
	return Node{
		Typ: nodeType,
		in:  in,
	}
}

func (n Node) Run() []interface{} {
	return n.Typ.Run(n.in)
}

func FastRun(nodeType NodeType, in []interface{}) []interface{} {
	return NewNode(nodeType, in).Run()
}
