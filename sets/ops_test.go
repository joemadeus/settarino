package sets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	tts := []struct {
		p1, p2 *PrimitiveSet
		okeys  []string
		name   string
	}{{
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("p2"), []string{}),
		okeys: []string{},
		name:  "empty_empty",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("p2"), []string{"piggy", "rowlf", "scooter"}),
		okeys: []string{},
		name:  "m_side_empty",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "gonzo", "kermit"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("p2"), []string{}),
		okeys: []string{},
		name:  "n_side_empty",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "gonzo", "kermit"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("p2"), []string{"rowlf", "scooter", "sweetums"}),
		okeys: []string{},
		name:  "disjoint",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"rowlf", "scooter", "sweetums"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("p2"), []string{"fozzie", "gonzo", "kermit"}),
		okeys: []string{},
		name:  "reverse_disjoint",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "gonzo", "rowlf", "sweetums"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("p2"), []string{"fozzie", "kermit", "piggy", "scooter"}),
		okeys: []string{},
		name:  "mixed_disjoint",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "kermit", "piggy", "scooter"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"chef", "gonzo", "rowlf", "sweetums"}),
		okeys: []string{},
		name:  "reverse_mixed_disjoint",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "fozzie", "gonzo", "kermit"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"chef", "fozzie", "piggy", "scooter"}),
		okeys: []string{"chef", "fozzie"},
		name:  "intersecting_start_start",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "fozzie", "gonzo", "kermit"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"gonzo", "kermit", "rowlf", "sweetums"}),
		okeys: []string{"gonzo", "kermit"},
		name:  "intersecting_end_start",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "fozzie", "gonzo", "kermit"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"beaker", "fozzie", "gonzo", "sweetums"}),
		okeys: []string{"fozzie", "gonzo"},
		name:  "intersecting_middle",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "fozzie", "harry", "kermit", "rowlf"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"beaker", "fozzie", "gonzo", "kermit", "sweetums"}),
		okeys: []string{"fozzie", "kermit"},
		name:  "intersecting_mixed",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "fozzie", "rowlf", "sweetums"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"beaker", "kermit", "rowlf", "sweetums"}),
		okeys: []string{"rowlf", "sweetums"},
		name:  "intersecting_end_end",
	}, {
		p1:    NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"chef", "fozzie", "rowlf", "sweetums"}),
		p2:    NewPrimitiveSet(time.Now(), CanonicalTag("poke"), []string{"animal", "beaker", "chef", "fozzie"}),
		okeys: []string{"chef", "fozzie"},
		name:  "intersecting_start_end",
	}}

	for _, tt := range tts {
		t.Run(tt.name+"_fast", func(t *testing.T) {
			o := FastIntersect(tt.p1, tt.p2.Elements())
			ds := eleSlice(o)

			assert.EqualValues(t, len(tt.okeys), len(ds))
			for i, e := range ds {
				assert.EqualValues(t, tt.okeys[i], e.Key)
			}
		})

		t.Run(tt.name+"_stream", func(t *testing.T) {
			o := StreamIntersect(tt.p1.Elements(), tt.p2.Elements())
			ds := eleSlice(o)

			assert.EqualValues(t, len(tt.okeys), len(ds))
			for i, e := range ds {
				assert.EqualValues(t, tt.okeys[i], e.Key)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	t.Run("empty_empty", func(t *testing.T) {
		m := toChan()
		n := toChan()
		o := Union(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("m_side_empty", func(t *testing.T) {
		m := toChan()
		n := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 3)
		assert.EqualValues(t, "piggy", ds[0].Key)
		assert.EqualValues(t, "rowlf", ds[1].Key)
		assert.EqualValues(t, "scooter", ds[2].Key)
	})

	t.Run("n_side_empty", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan()
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 3)
		assert.EqualValues(t, "fozzie", ds[0].Key)
		assert.EqualValues(t, "gonzo", ds[1].Key)
		assert.EqualValues(t, "kermit", ds[2].Key)
	})

	t.Run("disjoint", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 6)
		assert.EqualValues(t, "fozzie", ds[0].Key)
		assert.EqualValues(t, "gonzo", ds[1].Key)
		assert.EqualValues(t, "kermit", ds[2].Key)
		assert.EqualValues(t, "piggy", ds[3].Key)
		assert.EqualValues(t, "rowlf", ds[4].Key)
		assert.EqualValues(t, "scooter", ds[5].Key)
	})

	t.Run("disjoint_reverse", func(t *testing.T) {
		m := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"})
		n := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 6)
		assert.EqualValues(t, "fozzie", ds[0].Key)
		assert.EqualValues(t, "gonzo", ds[1].Key)
		assert.EqualValues(t, "kermit", ds[2].Key)
		assert.EqualValues(t, "piggy", ds[3].Key)
		assert.EqualValues(t, "rowlf", ds[4].Key)
		assert.EqualValues(t, "scooter", ds[5].Key)
	})

	t.Run("disjoint_mixed", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "animal"}, &Element{Key: "harry"}, &Element{Key: "scooter"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 6)
		assert.EqualValues(t, "animal", ds[0].Key)
		assert.EqualValues(t, "fozzie", ds[1].Key)
		assert.EqualValues(t, "gonzo", ds[2].Key)
		assert.EqualValues(t, "harry", ds[3].Key)
		assert.EqualValues(t, "kermit", ds[4].Key)
		assert.EqualValues(t, "scooter", ds[5].Key)
	})

	t.Run("disjoint_mixed_reversed", func(t *testing.T) {
		m := toChan(&Element{Key: "animal"}, &Element{Key: "harry"}, &Element{Key: "scooter"})
		n := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 6)
		assert.EqualValues(t, "animal", ds[0].Key)
		assert.EqualValues(t, "fozzie", ds[1].Key)
		assert.EqualValues(t, "gonzo", ds[2].Key)
		assert.EqualValues(t, "harry", ds[3].Key)
		assert.EqualValues(t, "kermit", ds[4].Key)
		assert.EqualValues(t, "scooter", ds[5].Key)
	})

	t.Run("equal_mixed", func(t *testing.T) {
		m := toChan(&Element{Key: "animal"}, &Element{Key: "fozzie"}, &Element{Key: "scooter"})
		n := toChan(&Element{Key: "fozzie"}, &Element{Key: "harry"}, &Element{Key: "kermit"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 5)
		assert.EqualValues(t, "animal", ds[0].Key)
		assert.EqualValues(t, "fozzie", ds[1].Key)
		assert.EqualValues(t, "harry", ds[2].Key)
		assert.EqualValues(t, "kermit", ds[3].Key)
		assert.EqualValues(t, "scooter", ds[4].Key)
	})

	t.Run("multi_equal_mixed", func(t *testing.T) {
		m := toChan(&Element{Key: "animal"}, &Element{Key: "fozzie"}, &Element{Key: "harry"}, &Element{Key: "scooter"})
		n := toChan(&Element{Key: "fozzie"}, &Element{Key: "harry"}, &Element{Key: "kermit"})
		o := Union(m, n)

		ds := eleSlice(o)
		assert.Len(t, ds, 5)
		assert.EqualValues(t, "animal", ds[0].Key)
		assert.EqualValues(t, "fozzie", ds[1].Key)
		assert.EqualValues(t, "harry", ds[2].Key)
		assert.EqualValues(t, "kermit", ds[3].Key)
		assert.EqualValues(t, "scooter", ds[4].Key)
	})
}

func toChan(eles ...*Element) chan *Element {
	out := make(chan *Element)
	go func() {
		defer close(out)
		for _, e := range eles {
			out <- e
		}
	}()

	return out
}

func eleSlice(es chan *Element) []*Element {
	out := make([]*Element, 0)
	for e := range es {
		out = append(out, e)
	}
	return out
}
