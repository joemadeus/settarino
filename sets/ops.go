package sets

import (
	"sync"
)

// FastIntersect performs an intersection on this and the provided Elements, but
// operates in O(n * log(n)) time instead of linear time. If 'poker' comes from a
// PrimitiveSet and its size is larger than ps, consider swapping the arguments' values
func FastIntersect(ps *PrimitiveSet, poker chan *Element) chan *Element {
	out := make(chan *Element)
	go func() {
		// locking here only because we don't get the locking in the Elements() func
		ps.RLock()
		defer ps.RUnlock()
		defer close(out)

		for p := range poker {
			if e := ps.member(p); e != nil {
				e.AddTags(p.Tags)
				out <- e
			}
		}
	}()

	return out
}

// StreamIntersect returns values on the output channel that exist in both input sets.
// Incoming values are assumed sorted within their chan, as are values on the output
// channel. Operates in linear time
func StreamIntersect(mD chan *Element, nD chan *Element) chan *Element {
	out := make(chan *Element)
	go func() {
		defer close(out)

		recvM := recv(mD)
		recvN := recv(nD)

		// provide nHi, mHi with an initial values so the loop has something to work on.
		// bonus: lets us exit quickly if either has no elements
		mHi, ok := recvM()
		if ok == false {
			drain(nD)
			return
		}

		nHi, ok := recvN()
		if ok == false {
			drain(mD)
			return
		}

		mMore, nMore := true, true
		for mMore && nMore {
			switch {
			case mHi.LessThan(nHi):
				mHi, mMore = recvM()

			case nHi.LessThan(mHi):
				nHi, nMore = recvN()

			default:
				// high water marks are equal, which means
				// this value is part of the intersection
				mHi.AddTags(nHi.Tags)
				out <- mHi
				mHi, mMore = recvM()
				nHi, nMore = recvN()
			}
		}
	}()

	return out
}

// Union performs a union operation on the two input channels and guarantees that the
// output is sorted and deduplicated. Incoming values are assumed sorted within their
// chan. Operates in linear time
func Union(mD chan *Element, nD chan *Element) chan *Element {
	out := make(chan *Element)
	go func() {
		defer close(out)

		recvM := recv(mD)
		recvN := recv(nD)

		mHi, mMore := recvM()
		nHi, nMore := recvN()

		for mMore || nMore {
			switch {
			case mHi.LessThan(nHi):
				out <- mHi
				mHi, mMore = recvM()

			case nHi.LessThan(mHi):
				out <- nHi
				nHi, nMore = recvN()

			case mHi.Equals(nHi):
				out <- mHi
				mHi, mMore = recvM()
				nHi, nMore = recvN()

			case mMore:
				out <- mHi
				mHi, mMore = recvM()

			case nMore:
				out <- nHi
				nHi, nMore = recvN()

			default:
				panic("assumption violation: comparison semantics have changed")
			}
		}
	}()

	return out
}

// Gather performs a union on the provided channels, but output is not guaranteed to be
// in sorted order, nor will it be deduplicated. This makes it unsuitable for input to
// other set operations, but is faster than Union if you need to combine the results of
// many operations, e.g., as the final step in a query. Operates in parallel on each
// input chan, in linear time
func Gather(in ...chan *Element) chan *Element {
	// impl note: the results here shouldn't be used as input to any Set, nor should they
	// be wrapped in a Set, since Set is defined with *ordered* elements and Gather
	// destroys order
	out := make(chan *Element)
	go func() {
		defer close(out)

		wg := sync.WaitGroup{}
		for _, c := range in {
			wg.Add(1)
			go func(c chan *Element) {
				defer wg.Done()
				for e := range c {
					out <- e
				}
			}(c)
		}

		wg.Wait()
	}()

	return out
}

func drain(in chan *Element) {
	for range in {
	}
}

// recv is a closure around a function that receives from the provided input channel,
// compares its latest value to a high water mark and panics if the elements are out of
// order. It returns the latest element and true if there is more to read, or a zero
// Element value and false if not
func recv(in chan *Element) func() (*Element, bool) {
	zero := &Element{}
	hi := &Element{}
	var cur *Element
	var ok bool

	return func() (*Element, bool) {
		cur, ok = <-in
		if ok {
			if cur.LessThan(hi) {
				// intentional panic
				panic("assumption violation: input is not sorted")
			}
			hi = cur
			return cur, ok
		} else {
			return zero, ok
		}
	}
}
