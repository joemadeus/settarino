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

	// Size returns the cardinality of the Elements in the set
	Size() int
}

type primitiveElement struct {
	key string
	tag Tag
}

// A PrimitiveSet is a SizedSet whose Elements are backed by a slice
type PrimitiveSet struct {
	sync.RWMutex
	tag        Tag
	primitives []primitiveElement
	lastupdate time.Time
}

// NewPrimitiveSet creates and returns a new PrimitiveSet with the provided Tag and
// keys
func NewPrimitiveSet(lup time.Time, tag Tag, keys []string) *PrimitiveSet {
	return &PrimitiveSet{
		primitives: createEles(tag, keys),
		tag:        tag,
		lastupdate: lup,
	}
}

func createEles(tag Tag, keys []string) []primitiveElement {
	eles := make([]primitiveElement, len(keys), len(keys))
	for i, k := range keys {
		eles[i] = primitiveElement{
			key: k,
			tag: tag,
		}
	}

	sort.Slice(eles, func(i, j int) bool {
		return lessthan(eles[i].key, eles[j].key)
	})

	return eles
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

		for i := 0; i < len(ps.primitives); i++ {
			// copy into an Element pointer. the original elements must not be modified in any way,
			// hence this assignment andthe internal type. this pattern is going to cause significant
			// GC pressure but it's easier than trying to manage the tags in other ways
			out <- &Element{
				Key:  ps.primitives[i].key,
				Tags: map[Tag]struct{}{ps.primitives[i].tag: {}},
			}
		}
	}()

	return out
}

// Member returns the Element in this Set with the provided key, or nil if it's not
// in this set
func (ps *PrimitiveSet) Member(poke string) *Element {
	i := sort.Search(len(ps.primitives), func(i int) bool {
		return lessthan(poke, ps.primitives[i].key)
	})

	i--

	if i >= 0 && i < len(ps.primitives) && ps.primitives[i].key == poke {
		// return a pointer to a copy of the element
		return &Element{
			Key:  ps.primitives[i].key,
			Tags: map[Tag]struct{}{ps.primitives[i].tag: {}},
		}
	}

	return nil
}

// LastUpdateTime returns the time of the most recent call to Reload or NewPrimitiveSet
func (ps *PrimitiveSet) LastUpdateTime() time.Time {
	ps.RLock()
	defer ps.RUnlock()

	return ps.lastupdate
}

// Reload replaces every element in the PrimitiveSet and sets the last update time to
// the provided time
func (ps *PrimitiveSet) Reload(lup time.Time, keys []string) {
	ps.Lock()
	defer ps.Unlock()

	ps.lastupdate = lup
	ps.primitives = createEles(ps.tag, keys)
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
