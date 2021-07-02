package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	t.Run("empty_empty", func(t *testing.T) {
		m := toChan()
		n := toChan()
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("m_side_empty", func(t *testing.T) {
		m := toChan()
		n := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"})
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("n_side_empty", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan()
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("disjoint", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"}, &Element{Key: "sweetums"})
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("reverse_disjoint", func(t *testing.T) {
		m := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"}, &Element{Key: "sweetums"})
		n := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("mixed_disjoint", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "gonzo"}, &Element{Key: "rowlf"}, &Element{Key: "sweetums"})
		n := toChan(&Element{Key: "fozzie"}, &Element{Key: "kermit"}, &Element{Key: "piggy"}, &Element{Key: "scooter"})
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("reverse_mixed_disjoint", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "kermit"}, &Element{Key: "piggy"}, &Element{Key: "scooter"})
		n := toChan(&Element{Key: "chef"}, &Element{Key: "gonzo"}, &Element{Key: "rowlf"}, &Element{Key: "sweetums"})
		o := StreamIntersect(m, n)

		d, ok := <-o
		assert.False(t, ok, "returned: %+v", d)
		assert.Nil(t, d)
	})

	t.Run("intersecting_start_start", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "piggy"}, &Element{Key: "scooter"})
		o := StreamIntersect(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "chef")
		assert.EqualValues(t, ds[1].Key, "fozzie")
	})

	t.Run("intersecting_end_start", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "gonzo"}, &Element{Key: "kermit"}, &Element{Key: "rowlf"}, &Element{Key: "sweetums"})
		o := StreamIntersect(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "gonzo")
		assert.EqualValues(t, ds[1].Key, "kermit")
	})

	t.Run("intersecting_middle", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "beaker"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "sweetums"})
		o := StreamIntersect(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "fozzie")
		assert.EqualValues(t, ds[1].Key, "gonzo")
	})

	t.Run("intersecting_mixed", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "harry"}, &Element{Key: "kermit"}, &Element{Key: "rowlf"})
		n := toChan(&Element{Key: "beaker"}, &Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"}, &Element{Key: "sweetums"})
		o := StreamIntersect(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "fozzie")
		assert.EqualValues(t, ds[1].Key, "kermit")
	})

	t.Run("intersecting_end_end", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "rowlf"}, &Element{Key: "sweetums"})
		n := toChan(&Element{Key: "beaker"}, &Element{Key: "kermit"}, &Element{Key: "rowlf"}, &Element{Key: "sweetums"})
		o := StreamIntersect(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "rowlf")
		assert.EqualValues(t, ds[1].Key, "sweetums")
	})

	t.Run("intersecting_start_end", func(t *testing.T) {
		m := toChan(&Element{Key: "chef"}, &Element{Key: "fozzie"}, &Element{Key: "rowlf"}, &Element{Key: "sweetums"})
		n := toChan(&Element{Key: "animal"}, &Element{Key: "beaker"}, &Element{Key: "chef"}, &Element{Key: "fozzie"})
		o := StreamIntersect(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "chef")
		assert.EqualValues(t, ds[1].Key, "fozzie")
	})
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

		ds := toSlice(o)
		assert.Len(t, ds, 3)
		assert.EqualValues(t, "piggy", ds[0].Key)
		assert.EqualValues(t, "rowlf", ds[1].Key)
		assert.EqualValues(t, "scooter", ds[2].Key)
	})

	t.Run("n_side_empty", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan()
		o := Union(m, n)

		ds := toSlice(o)
		assert.Len(t, ds, 3)
		assert.EqualValues(t, "fozzie", ds[0].Key)
		assert.EqualValues(t, "gonzo", ds[1].Key)
		assert.EqualValues(t, "kermit", ds[2].Key)
	})

	t.Run("disjoint", func(t *testing.T) {
		m := toChan(&Element{Key: "fozzie"}, &Element{Key: "gonzo"}, &Element{Key: "kermit"})
		n := toChan(&Element{Key: "piggy"}, &Element{Key: "rowlf"}, &Element{Key: "scooter"})
		o := Union(m, n)

		ds := toSlice(o)
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

		ds := toSlice(o)
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

		ds := toSlice(o)
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

		ds := toSlice(o)
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

		ds := toSlice(o)
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

		ds := toSlice(o)
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

func toSlice(es chan *Element) []*Element {
	out := make([]*Element, 0)
	for e := range es {
		out = append(out, e)
	}
	return out
}
