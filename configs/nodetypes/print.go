package nodetypes

import (
	"fmt"
	"notabug.org/bosko/erigami/pkg/node"
)

func Print() node.NodeType {
	return node.NodeType{
		Name: "Print",
		In:   node.Rtype([]interface{}(nil)),
		Out:  nil,
		Vrun: func(value []interface{}) []interface{} {
			for _, v := range value {
				fmt.Println(v)
			}
			return nil
		},
	}
}
