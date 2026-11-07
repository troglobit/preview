package details

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Transformer replaces HTML blocks starting with <details> by Details nodes
type Transformer struct {
	idPrefix string
	counter  int
}

// Transform walks the AST and replaces HTML blocks containing <details> tags
func (t *Transformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	source := reader.Source()
	var blocks []*ast.HTMLBlock

	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if htmlBlock, ok := n.(*ast.HTMLBlock); ok && entering {
			content := bytes.TrimSpace(htmlBlock.Lines().Value(source))
			if bytes.HasPrefix(content, []byte("<details")) {
				blocks = append(blocks, htmlBlock)
			}
		}
		return ast.WalkContinue, nil
	})

	for _, htmlBlock := range blocks {
		t.counter++
		details := &Details{ID: t.generateID(htmlBlock.Lines().Value(source))}
		details.SetLines(htmlBlock.Lines())
		htmlBlock.Parent().ReplaceChild(htmlBlock.Parent(), htmlBlock, details)
	}
}

// generateID creates a unique ID for a details element based on its content
func (t *Transformer) generateID(content []byte) string {
	hash := sha256.Sum256(content)
	hashStr := hex.EncodeToString(hash[:])[:12]
	return fmt.Sprintf("%s%d-%s", t.idPrefix, t.counter, hashStr)
}

// NewTransformer creates a new Transformer
func NewTransformer(idPrefix string) *Transformer {
	return &Transformer{
		idPrefix: idPrefix,
		counter:  0,
	}
}
