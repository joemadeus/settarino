package catalog

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nytimes/settarino/sets"
)

func TestRoundTrip(t *testing.T) {
	fozzie, gonzo, rawlf := "fozzie", "gonzo", "rawlf"
	wps := sets.NewPrimitiveSet(
		time.Now(),
		sets.CanonicalTag("MUPN"),
		[]string{fozzie, gonzo, rawlf})

	w := bytes.Buffer{}
	if err := PersistSet(&w, wps); err != nil {
		t.Fatalf("did not persist the test set: %+v", err)
	}

	r := bytes.NewBuffer(w.Bytes())
	rps, err := LoadSet(r)
	if err != nil {
		t.Fatalf("did not reload the test set: %+v", err)
	}

	assert.NotNil(t, rps)
	// second resolution is fine. the protobuf code seems to mess things up
	// at finer resolutions
	assert.EqualValues(t, wps.LastUpdateTime().Unix(), rps.LastUpdateTime().Unix())
	assert.EqualValues(t, wps.Tag(), rps.Tag())

	reles := make([]*sets.Element, 0)
	for re := range rps.Elements() {
		reles = append(reles, re)
	}
	assert.Len(t, reles, 3)
	assert.EqualValues(t, fozzie, reles[0].Key)
	assert.EqualValues(t, gonzo, reles[1].Key)
	assert.EqualValues(t, rawlf, reles[2].Key)
}
