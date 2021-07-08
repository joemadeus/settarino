package sets

// Element is the type on which set operations operate, using the Key field as their
// basis. Tags are references to Set identifiers; they are appended whenever an
// Element instance is "selected" by an intersection
type Element struct {
	Key  string
	Tags map[Tag]struct{}
}

// Equals returns true if neither key is the empty string and both keys are equal
func (e *Element) Equals(e2 *Element) bool {
	return e.Key != "" && e.Key == e2.Key
}

// LessThan returns true if the this Element's Key is lexicographically less than the
// provided Element's Key. A zero Key is never less than any non-zero Key, and two zero
// Keys are never LessThan
func (e *Element) LessThan(dd *Element) bool {
	return lessthan(e.Key, dd.Key)
}

func lessthan(k1, k2 string) bool {
	return k1 != "" && k2 != "" && k1 < k2
}

func (e *Element) String() string {
	return e.Key
}

func (e *Element) AddTags(ts map[Tag]struct{}) {
	for t, s := range ts {
		e.Tags[t] = s
	}
}

// An Install Element stores an installation token
type Install Element

// An Email Element stores a newsletter subscription email address
type Email Element
