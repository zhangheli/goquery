package goquery

import (
	"code.google.com/p/cascadia"
	"code.google.com/p/go.net/html"
)

// Is checks the current matched set of elements against a selector and
// returns true if at least one of these elements matches.
func (s *Selection) Is(selector string) bool {
	if len(s.Nodes) > 0 {
		// Attempt a match with the selector
		cs := cascadia.MustCompile(selector)
		if len(s.Nodes) == 1 {
			return cs.Match(s.Nodes[0])
		}
		return len(cs.Filter(s.Nodes)) > 0
	}

	return false
}

// IsFunction checks the current matched set of elements against a predicate and
// returns true if at least one of these elements matches.
func (s *Selection) IsFunction(f func(int, *Selection) bool) bool {
	return s.FilterFunction(f).Length() > 0
}

// IsSelection checks the current matched set of elements against a Selection object
// and returns true if at least one of these elements matches.
func (s *Selection) IsSelection(sel *Selection) bool {
	return s.FilterSelection(sel).Length() > 0
}

// IsNodes checks the current matched set of elements against the specified nodes
// and returns true if at least one of these elements matches.
func (s *Selection) IsNodes(nodes ...*html.Node) bool {
	return s.FilterNodes(nodes...).Length() > 0
}

// Contains returns true if the specified Node is within,
// at any depth, one of the nodes in the Selection object.
// It is NOT inclusive, to behave like jQuery's implementation, and
// unlike Javascript's .contains, so if the contained
// node is itself in the selection, it returns false.
func (s *Selection) Contains(n *html.Node) bool {
	return sliceContains(s.Nodes, n)
}
