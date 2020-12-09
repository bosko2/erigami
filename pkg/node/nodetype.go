package node

import "reflect"

//Node type definition( restrictions and data)
type NodeType struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type
	Vrun func(value []interface{}) []interface{}
}

func (n *NodeType) Run(in []interface{}) []interface{} {
	return n.Vrun(in)
}

func Rtype(t ...interface{}) (a []reflect.Type) {
	for _, it := range t {
		a = append(a, reflect.TypeOf(it))
	}
	return
}
