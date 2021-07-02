package sets

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Set interface {
	// Elements returns a channel over the ordered elements in this Set
	Elements() chan *Element
}

// SizedSet adds cardinality retrieval to the Set type
type SizedSet interface {
	Set

	// Size returns the cardinality of the Elements in this set
	Size() int
}

// A PrimitiveSet is a SizedSet whose Elements are backed by a slice
type PrimitiveSet struct {
	sync.RWMutex
	tag        Tag
	primitives []*Element
	lastupdate time.Time
}

func NewPrimitiveSet(lup time.Time, eles []*Element) *PrimitiveSet {
	return &PrimitiveSet{
		primitives: eles,
		lastupdate: lup,
	}
}

func (ps *PrimitiveSet) Intersect(o Set) *StreamingSet {
	var out chan *Element
	switch other := o.(type) {
	case *PrimitiveSet:
		if other.Size() < ps.Size() {
			out = FastIntersect(ps, other.Elements())
		}
		out = FastIntersect(other, ps.Elements())

	case *StreamingSet:
		out = FastIntersect(ps, other.Elements())

	default:
		panic(fmt.Sprintf("assumption violation: %T isn't handled by PrimitiveSet", other))
	}

	return &StreamingSet{elements: out}
}

func (ps *PrimitiveSet) Union(o Set) *StreamingSet {
	return &StreamingSet{
		elements: Union(ps.Elements(), o.Elements()),
	}
}

func (ps *PrimitiveSet) Elements() chan *Element {
	out := make(chan *Element)
	go func() {
		ps.RLock()
		defer ps.RUnlock()
		defer close(out)
		for _, e := range ps.primitives {
			out <- e
		}
	}()

	return out
}

// Member returns the Element in this Set with the provided key, or nil if it's not
// in this set
func (ps *PrimitiveSet) Member(key string) *Element {
	return ps.member(&Element{Key: key})
}

func (ps *PrimitiveSet) member(poke *Element) *Element {
	i := sort.Search(len(ps.primitives), func(i int) bool {
		return ps.primitives[i].LessThan(poke) == false
	})

	if i < len(ps.primitives) && ps.primitives[i].Equals(poke) {
		return ps.primitives[i]
	}

	return nil
}

func (ps *PrimitiveSet) LastUpdateTime() time.Time {
	ps.RLock()
	defer ps.RUnlock()

	return ps.lastupdate
}

func (ps *PrimitiveSet) Reload(lup time.Time, eles []*Element) {
	ps.Lock()
	defer ps.Unlock()

	ps.lastupdate = lup
	ps.primitives = eles
}

func (ps *PrimitiveSet) Size() int {
	ps.RLock()
	defer ps.RUnlock()

	return len(ps.primitives)
}

func (ps *PrimitiveSet) Tag() Tag {
	return ps.tag
}

// A StreamingSet is a Set whose Elements are provided by a channel
type StreamingSet struct {
	elements chan *Element
}

func (ss *StreamingSet) Intersect(o Set) *StreamingSet {
	switch other := o.(type) {
	case *PrimitiveSet:
		return &StreamingSet{elements: FastIntersect(other, ss.Elements())}

	case *StreamingSet:
		return &StreamingSet{elements: StreamIntersect(ss.Elements(), other.Elements())}

	default:
		panic(fmt.Sprintf("assumption violation: %T isn't handled by PrimitiveSet", other))
	}
}

func (ss *StreamingSet) Union(o Set) *StreamingSet {
	return &StreamingSet{
		elements: Union(ss.Elements(), o.Elements()),
	}
}

func (ss *StreamingSet) Elements() chan *Element {
	return ss.elements
}

func tagSlice(eles []*Element, ts ...Tag) {
	for _, e := range eles {
		if e.Tags == nil {
			e.Tags = make(map[Tag]struct{})
		}

		for _, t := range ts {
			e.Tags[t] = struct{}{}
		}
	}
}
