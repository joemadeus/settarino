package query

import (
	"sync"

	"github.com/nytimes/settarino/catalog"
	"github.com/nytimes/settarino/sets"
)

// RoughQuery is a programmatic query that returns Elements with a BNA channel
// and are of type newsfusion, for only RI
func RoughQuery(cat *catalog.Catalog) (chan *sets.Element, error) {
	bnaSet, err := cat.PrimitiveSet(sets.CanonicalTag("chan:PUBN"))
	if err != nil {
		return nil, err
	}

	nfSet, err := cat.PrimitiveSet(sets.CanonicalTag("install:newsfusion"))
	if err != nil {
		return nil, err
	}

	riSet, err := cat.PrimitiveSet(sets.CanonicalTag("geo:US_RI"))
	if err != nil {
		return nil, err
	}

	return bnaSet.Intersect(riSet.Intersect(nfSet)).Elements(), nil
}

// Membership returns the slice of sets.Tag to which the provided key belongs
func Membership(cat *catalog.Catalog, key string) []sets.Tag {
	accumulate := make(chan sets.Tag)
	defer close(accumulate)

	wg := sync.WaitGroup{}
	for _, t := range cat.Tags() {
		wg.Add(1)
		go func(tag sets.Tag) {
			defer wg.Done()
			set, err := cat.PrimitiveSet(tag)
			if err != nil {
				// TODO: log this
				return
			}

			if set.Member(key) != nil {
				accumulate <- tag
			}
		}(t)
	}

	tags := make([]sets.Tag, 0)

	go func() {
		for t := range accumulate {
			tags = append(tags, t)
		}
	}()

	wg.Wait()

	return tags
}
