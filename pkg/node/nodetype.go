package node

import "reflect"

//Node type definition
type NodeType struct {
	In   []reflect.Type
	Out  []reflect.Type
	Vrun func(value []interface{}) []interface{}
}

//Runs the nodetypes function Vrun
func (n *NodeType) Run(in []interface{}) []interface{} {
	return n.Vrun(in)
}

//gets types of interfaces t
func Rtype(t ...interface{}) (a []reflect.Type) {
	for _, it := range t {
		a = append(a, reflect.TypeOf(it))
	}
	return
}
