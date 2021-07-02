package sets

import (
	"strings"
	"sync"
)

type CanonicalString *string

// StringCanon maintains pointers for a set of canonical strings known to a runtime. This
// type is not suitable for Unicode strings
type StringCanon struct {
	sync.RWMutex
	ss map[string]CanonicalString
}

func (sc *StringCanon) get(s string) CanonicalString {
	t := sc.read(s)
	if t != nil {
		return t
	}

	return sc.write(s)
}

func (sc *StringCanon) read(s string) CanonicalString {
	sc.RLock()
	defer sc.RUnlock()

	t, ok := sc.ss[s]
	if ok {
		return t
	}

	return nil
}

func (sc *StringCanon) write(s string) CanonicalString {
	sc.Lock()
	defer sc.Unlock()

	if t, ok := sc.ss[s]; ok {
		return t
	}

	sptr := &s
	sc.ss[s] = sptr
	return sptr
}

// A Tag defines a user-specified property of Sets and Elements
type Tag CanonicalString

// tagcanon contains all the canonical Tags in use by this runtime
var tagcanon StringCanon = StringCanon{ss: make(map[string]CanonicalString)}

// CanonicalTag returns a pointer to the canonical Tag for the provided string. A
// Tag's canonical string is always normalized to lowercase
func CanonicalTag(s string) Tag {
	return Tag(tagcanon.get(strings.ToLower(s)))
}
