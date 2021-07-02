package sets

// Element is the type on which set operations operate, using the Key field as their
// basis. Tags are references to Set identifiers; they are appended whenever an
// Element instance is "selected" by an intersection
type Element struct {
	Key  string
	Tags map[Tag]struct{}
}

func (d *Element) Equals(dd *Element) bool {
	return d.Key != "" && dd.Key != "" && d.Key == dd.Key
}

// LessThan returns true if the this Element's Key is lexicographically less than the
// provided Element's Key. A zero Key is never less than any non-zero Key, and two zero
// Keys are never LessThan
func (d *Element) LessThan(dd *Element) bool {
	return d.Key != "" && dd.Key != "" && d.Key < dd.Key
}

func (d *Element) String() string {
	return d.Key
}

func (d *Element) AddTags(ts map[Tag]struct{}) {
	for t, s := range ts {
		d.Tags[t] = s
	}
}

// An Install Element stores an installation token
type Install Element

// An Email Element stores a newsletter subscription address
type Email Element
