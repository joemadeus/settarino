package catalog

import (
	"fmt"
	"sync"

	"github.com/nytimes/settarino/sets"
)

type NoSuchTagError struct {
	tag sets.Tag
}

func (n NoSuchTagError) Error() string {
	return fmt.Sprintf("no tag %s known", *n.tag)
}

type Catalog struct {
	sync.RWMutex
	sets map[sets.Tag]*sets.PrimitiveSet
}

func (c *Catalog) AddPrimitiveSet(tag sets.Tag, set *sets.PrimitiveSet) error {
	c.Lock()
	defer c.Unlock()

	if k, _ := c.sets[tag]; k != nil {
		return fmt.Errorf("set with tag %v is already known in the catalog", tag)
	}

	c.sets[tag] = set
	return nil
}

// PrimitiveSet returns the sets.PrimitiveSet associated with the given sets.Tag, or
// NoSuchTagError if the Tag is unknown
func (c *Catalog) PrimitiveSet(tag sets.Tag) (*sets.PrimitiveSet, *NoSuchTagError) {
	c.RLock()
	defer c.RUnlock()

	s, ok := c.sets[tag]
	if ok == false {
		return nil, &NoSuchTagError{tag: tag}
	}

	return s, nil
}

// Tags returns a slice of sets.Tag known to this Catalog
func (c *Catalog) Tags() []sets.Tag {
	c.RLock()
	defer c.RUnlock()

	out := make([]sets.Tag, 0, len(c.sets))
	for t, _ := range c.sets {
		out = append(out, t)
	}

	return out
}
