package details

import "github.com/yuin/goldmark/ast"

// KindDetails is the kind for Details nodes
var KindDetails = ast.NewNodeKind("Details")

// Details is a raw HTML block starting with <details>, with the id to inject
// into its opening tag.
type Details struct {
	ast.BaseBlock
	ID string
}

// Kind returns the kind of this node
func (n *Details) Kind() ast.NodeKind {
	return KindDetails
}

// Dump dumps the details node
func (n *Details) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, map[string]string{"ID": n.ID}, nil)
}
