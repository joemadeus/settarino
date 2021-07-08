package sets

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ContainsTag(t *testing.T, tags map[Tag]struct{}, contains Tag, msgAndArgs ...interface{}) bool {
	for tag := range tags {
		if tag == contains {
			return true
		}
	}

	return assert.Fail(t, fmt.Sprintf("%#v does not contain %#v", tags, contains), msgAndArgs...)
}

func TestNewPrimitiveSet(t *testing.T) {
	p1Tag := CanonicalTag("p1")
	p1Keys := []string{"fozzie", "kermit", "gonzo"}
	p1 := NewPrimitiveSet(time.Now(), p1Tag, p1Keys)

	assert.EqualValues(t, "fozzie", p1.primitives[0].key)
	assert.Equal(t, p1Tag, p1.primitives[0].tag)
	assert.EqualValues(t, "gonzo", p1.primitives[1].key)
	assert.Equal(t, p1Tag, p1.primitives[1].tag)
	assert.EqualValues(t, "kermit", p1.primitives[2].key)
	assert.Equal(t, p1Tag, p1.primitives[2].tag)
}

func TestPrimitiveSet_Elements(t *testing.T) {
	p1Tag := CanonicalTag("p1")
	p1Keys := []string{"fozzie", "kermit", "gonzo"}
	p1 := NewPrimitiveSet(time.Now(), p1Tag, p1Keys)

	ps := eleSlice(p1.Elements())
	assert.EqualValues(t, "fozzie", ps[0].Key)
	ContainsTag(t, ps[0].Tags, p1Tag)
	assert.EqualValues(t, "gonzo", ps[1].Key)
	ContainsTag(t, ps[1].Tags, p1Tag)
	assert.EqualValues(t, "kermit", ps[2].Key)
	ContainsTag(t, ps[2].Tags, p1Tag)
}

func TestPrimitiveSet_Member(t *testing.T) {
	t.Run("no_member", func(t *testing.T) {
		seek := "animal"
		p1 := NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "gonzo", "kermit"})
		assert.Nil(t, p1.Member(seek))
	})

	t.Run("start_member", func(t *testing.T) {
		seek := "fozzie"
		p1 := NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "gonzo", "kermit"})
		m := p1.Member(seek)
		assert.NotNil(t, m)
		assert.EqualValues(t, m.Key, seek)
	})

	t.Run("middle_member", func(t *testing.T) {
		seek := "gonzo"
		p1 := NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "gonzo", "kermit"})
		m := p1.Member(seek)
		assert.NotNil(t, m)
		assert.EqualValues(t, m.Key, seek)
	})

	t.Run("end_member", func(t *testing.T) {
		seek := "kermit"
		p1 := NewPrimitiveSet(time.Now(), CanonicalTag("p1"), []string{"fozzie", "gonzo", "kermit"})
		m := p1.Member(seek)
		assert.NotNil(t, m)
		assert.EqualValues(t, m.Key, seek)
	})
}

func TestPrimitiveSet_Intersect(t *testing.T) {
	t.Run("primitive_primitive", func(t *testing.T) {
		p1Tag, p2Tag := CanonicalTag("p1"), CanonicalTag("p2")
		p1 := NewPrimitiveSet(time.Now(), p1Tag, []string{"fozzie", "gonzo", "kermit"})
		p2 := NewPrimitiveSet(time.Now(), p2Tag, []string{"gonzo", "kermit", "scooter"})

		o := p1.Intersect(p2)

		ds := eleSlice(o.Elements())
		assert.Equal(t, len(ds), 2)
		assert.EqualValues(t, ds[0].Key, "gonzo")
		assert.EqualValues(t, ds[1].Key, "kermit")

		for _, d := range ds {
			assert.Len(t, d.Tags, 2)
			ContainsTag(t, d.Tags, p1Tag, "%v", d)
			ContainsTag(t, d.Tags, p2Tag, "%v", d)
		}
	})
}
