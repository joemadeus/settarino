package query

type SubQ interface {
	// TODO
}

// Segment is an attribute of a set of installations
type Segment string

// subQ is a subquery, which can identify Segments and other SubQs. A SubQ's
// behavior applies to all its Segments and to the output of each of its SubQs
type subQ struct {
	Segs  []Segment
	Subqs []SubQ
}

// AudQ is the root of an audience query
type AudQ subQ

// JoinQ provides the intersection of its Segments and SubQs
type JoinQ subQ

// AllQ provides the union of its Segments and SubQs
type AllQ subQ

// NotQ provides the negation of its Segments and SubQs
type NotQ subQ
