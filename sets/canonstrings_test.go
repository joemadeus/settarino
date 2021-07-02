package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCanon_get(t *testing.T) {
	canon := &StringCanon{
		ss: make(map[string]CanonicalString),
	}
	assert.Nil(t, canon.read("ri"))
	cs := canon.get("ri")
	assert.EqualValues(t, "ri", *cs)

	cs2 := canon.get("ri")
	assert.Same(t, cs, cs2)
}

func TestStringCanon_rw(t *testing.T) {
	t.Run("nobody_home", func(t *testing.T) {
		canon := &StringCanon{
			ss: make(map[string]CanonicalString),
		}
		assert.Nil(t, canon.read("ri"))
	})

	t.Run("why_hello_there", func(t *testing.T) {
		canon := &StringCanon{
			ss: make(map[string]CanonicalString),
		}
		sri, sct := "ri", "ct"

		assert.Nil(t, canon.read(sri))
		assert.Nil(t, canon.read(sct))

		csri := canon.write(sri)
		csct := canon.write(sct)
		assert.EqualValues(t, sri, *csri)
		assert.EqualValues(t, sct, *csct)
		assert.NotEqualValues(t, *csri, *csct)
		assert.NotSame(t, csri, csct)

		csri2 := canon.read("ri")
		csct2 := canon.read("ct")
		assert.EqualValues(t, sri, *csri2)
		assert.EqualValues(t, sct, *csct2)
		assert.Same(t, csri, csri2)
		assert.Same(t, csct, csct2)
		assert.NotEqualValues(t, *csri2, *csct2)
		assert.NotSame(t, csri2, csct2)
	})
}
