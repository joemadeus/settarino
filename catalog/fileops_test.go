package catalog

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nytimes/settarino/sets"
)

func TestRoundTrip(t *testing.T) {
	fozzie := sets.Element{
		Key:     "fozzie",
		Payload: []byte{0x1},
	}

	gonzo := sets.Element{
		Key:     "gonzo",
		Payload: []byte{0x0},
	}

	rawlf := sets.Element{
		Key:     "rawlf",
		Payload: []byte{0x1, 0x1},
	}

	wlup := time.Now()
	wsid := "MUPN"
	weles := []sets.Element{fozzie, gonzo, rawlf}
	wset := sets.NewPrimitiveSet(wlup, weles)

	w := bytes.Buffer{}
	if err := PersistSet(&w, wsid, wset); err != nil {
		t.Fatalf("did not persist the test set: %+v", err)
	}

	r := bytes.NewBuffer(w.Bytes())
	rsid, rset, err := LoadSet(r)
	if err != nil {
		t.Fatalf("did not reload the test set: %+v", err)
	}

	assert.NotNil(t, rset)

	reles := make([]sets.Element, 0)
	for re := range rset.Elements() {
		reles = append(reles, re)
	}

	rlup := rset.LastUpdateTime()
	assert.True(t, wlup.Equal(rlup), "wanted %+v, got %+v", wlup, rlup)
	assert.EqualValues(t, wsid, rsid)
	assert.Len(t, reles, 3)
	assert.EqualValues(t, fozzie, reles[0])
	assert.EqualValues(t, gonzo, reles[1])
	assert.EqualValues(t, rawlf, reles[2])
}
