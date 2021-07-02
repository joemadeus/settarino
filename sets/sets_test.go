package sets

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrimitiveSet_TagSlice(t *testing.T) {
	tag := CanonicalTag("p1")
	eles := []*Element{{Key: "fozzie"}, {Key: "gonzo"}, {Key: "kermit"}}
	tagSlice(eles, tag)

	assert.Len(t, eles, 3)
	for _, o := range eles {
		assert.Len(t, o.Tags, 1)
		assert.Contains(t, o.Tags, tag)
	}
}

func ContainsTag(t *testing.T, tags map[Tag]struct{}, contains Tag, msgAndArgs ...interface{}) bool {
	for tag := range tags {
		if tag == contains {
			return true
		}
	}

	return assert.Fail(t, fmt.Sprintf("%#v does not contain %#v", tags, contains), msgAndArgs...)
}

func TestPrimitiveSet_Intersect(t *testing.T) {
	t.Run("primitive_primitive", func(t *testing.T) {
		p1Tag, p2Tag := CanonicalTag("p1"), CanonicalTag("p2")
		p1Eles := []*Element{{Key: "fozzie"}, {Key: "gonzo"}, {Key: "kermit"}}
		p2Eles := []*Element{{Key: "gonzo"}, {Key: "kermit"}, {Key: "scooter"}}
		tagSlice(p1Eles, p1Tag)
		tagSlice(p2Eles, p2Tag)

		p1 := NewPrimitiveSet(time.Now(), p1Eles)
		p2 := NewPrimitiveSet(time.Now(), p2Eles)
		o := p1.Intersect(p2)

		ds := toSlice(o.Elements())
		assert.Len(t, ds, 2)
		assert.EqualValues(t, ds[0].Key, "gonzo")
		assert.EqualValues(t, ds[1].Key, "kermit")

		for _, d := range ds {
			assert.Len(t, d.Tags, 2)
			ContainsTag(t, d.Tags, p1Tag, "%v", d)
			ContainsTag(t, d.Tags, p2Tag, "%v", d)
		}
	})
}
