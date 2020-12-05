package nodetyp

//Node type definition( restrictions and data)
type NodeType struct {
	Nin   []*NodeType
	Name  string
	Nout  []*NodeType
	Tpath string
	Mode  string
}
