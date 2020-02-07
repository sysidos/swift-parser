// generated by Textmapper; DO NOT EDIT

package ast

import (
	"sort"
	"strings"
	"swift-grammer/js"
	"swift-grammer/js/selector"
)

// Tree is a parse tree for some content.
type Tree struct {
	path    string
	content string
	lines   []int
	root    *Node
}

// NewTree creates a new Tree.
func newTree(path, content string) *Tree {
	return &Tree{path: path, content: content, lines: lineOffsets(content)}
}

// Path returns the location of the parsed content (if any).
func (t *Tree) Path() string {
	return t.path
}

// Root returns the root node of this tree.
func (t *Tree) Root() *Node {
	return t.root
}

// Text returns the parsed content.
func (t *Tree) Text() string {
	return t.content
}

// Node is an AST node.
type Node struct {
	t          js.NodeType
	offset     int
	endoffset  int
	parent     *Node
	next       *Node
	firstChild *Node
	tree       *Tree
}

// IsValid helps to detect non-existing nodes.
func (n *Node) IsValid() bool {
	return n != nil
}

// Type returns
func (n *Node) Type() js.NodeType {
	if n == nil {
		return js.NoType
	}
	return n.t
}

// Offset returns the start offset of the node.
func (n *Node) Offset() int {
	if n == nil {
		return 0
	}
	return n.offset
}

// Endoffset returns the end offset of the node.
func (n *Node) Endoffset() int {
	if n == nil {
		return 0
	}
	return n.endoffset
}

// LineColumn returns the start position of the nodes as 1-based line and column.
func (n *Node) LineColumn() (int, int) {
	if n == nil {
		return 1, 1
	}
	lines := n.tree.lines
	offset := n.offset
	line := sort.Search(len(lines), func(i int) bool { return lines[i] > offset }) - 1
	return line + 1, offset - lines[line] + 1
}

// Child returns the first child node matching a given selector.
func (n *Node) Child(sel selector.Selector) *Node {
	if n == nil {
		return nil
	}
	for c := n.firstChild; c != nil; c = c.next {
		if sel(c.t) {
			return c
		}
	}
	return nil
}

// Children returns all child nodes matching a given selector.
func (n *Node) Children(sel selector.Selector) []*Node {
	if n == nil {
		return nil
	}
	var ret []*Node
	for c := n.firstChild; c != nil; c = c.next {
		if sel(c.t) {
			ret = append(ret, c)
		}
	}
	return ret
}

// Next returns the first node among right siblings of this node matching a given selector.
func (n *Node) Next(sel selector.Selector) *Node {
	if n == nil {
		return nil
	}
	for c := n.next; c != nil; c = c.next {
		if sel(c.t) {
			return c
		}
	}
	return nil
}

// NextAll return all right siblings of this node matching a given selector.
func (n *Node) NextAll(sel selector.Selector) []*Node {
	if n == nil {
		return nil
	}
	var ret []*Node
	for c := n.next; c != nil; c = c.next {
		if sel(c.t) {
			ret = append(ret, c)
		}
	}
	return ret
}

// Text returns the text of the node.
func (n *Node) Text() string {
	if n == nil {
		return ""
	}
	return n.tree.content[n.offset:n.endoffset]
}

func lineOffsets(str string) []int {
	var lines = make([]int, 1, 128)

	var off int
	for {
		i := strings.IndexByte(str[off:], '\n')
		if i == -1 {
			break
		}
		off += i + 1
		lines = append(lines, off)
	}
	return lines
}
