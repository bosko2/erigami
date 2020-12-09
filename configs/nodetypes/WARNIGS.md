This is a file witch explains rules of defining new node type.
1. Every type has its own `name.go` inside `configs'/nodetype` and `name.tmpl` inside `web/nodes`

name.go:

```go
//put name of a function 
package nodetypes

import (
	"notabug.org/bosko/erigami/pkg/node"
	"reflect"
)

//explain what it does here in one line
types["Name"]= node.NodeType {
	return node.NodeType{
		In:   node.Rtype((/*type for input goes here"*/int)(nil),(int/*secons type*/)(nil)/*...*/)),
		Out:  node.Rtype(/*same as input*/), 
		/*In and Out can be nil*/
		Vrun: func(value []interface{}) []interface{
		    //code goes here	
			return //always return something even if Out is nil(then return nil)
        } ,
	}
	/*Name and function name have to be uppercase*/
}
```

