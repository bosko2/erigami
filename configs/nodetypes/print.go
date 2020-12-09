package nodetypes

types["Print"]= node.NodeType {
	return node.NodeType{
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
