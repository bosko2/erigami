package node

/*
* TODO: TEST!
 */

import (
	"bytes"
	"notabug.org/bosko/erigami/pkg/nodetyp"
	"text/template"
)

//type for a node(command)
type Node struct {
	Typ  nodetyp.NodeType
	Nin  *[]Node
	in   interface{}
	nout *[]Node
	out  interface{}
	Code string
}

//adds output if it is valid
func (this *Node) AddOutput(o int, i int, n *Node) bool {
	if this.Typ.Nout[o].Name == n.Typ.Name && n.TestInput(i, this) {
		a := *(this.nout)
		a[o] = *n
		this.nout = &a
		b := *(n.Nin)
		b[i] = *this
		n.Nin = &b
		return true
	}
	return false
}

//tests if input is valid
func (this *Node) TestInput(i int, n *Node) bool {
	if this.Typ.Nin[i].Name == n.Typ.Name {
		return true
	}
	return false
}

func (this *Node) GenCode() {
	t, err := template.ParseFiles(this.Typ.Tpath)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = t.Execute(&b, *this)
	if err != nil {
		panic(err)
	}
	this.Code = b.String()
}
