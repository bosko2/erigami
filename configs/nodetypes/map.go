package nodetypes

import (
	"dev.azure.com/markobosnjak20/erigami/_git/erigami/pkg/node"
	"fmt"
)

var Types = map[string]node.NodeType{
	"Print": node.NodeType{
		In:  node.Rtype(""),
		Out: nil,
		Vrun: func(value []interface{}) []interface{} {
			for _, v := range value {
				fmt.Println(v)
			}
			return nil
		},
	},
	"AskForString": node.NodeType{
		In:  nil,
		Out: node.Rtype(""),
		Vrun: func(value []interface{}) []interface{} {
			/*var s string
			fmt.Scanf("%s",&s)*/
			return []interface{}{"w"}
		},
	},
}
