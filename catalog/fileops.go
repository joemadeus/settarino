package catalog

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nytimes/settarino/sets"
)

// LoadSet creates a new sets.PrimitiveSet from the protobuf encoded bytes provided by
// the given io.Reader
func LoadSet(r io.Reader) (*sets.PrimitiveSet, error) {
	slurp, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	pbwrapper := &ElementFile{}
	if err := proto.Unmarshal(slurp, pbwrapper); err != nil {
		return nil, err
	}

	pbdata := pbwrapper.GetRstyle()
	wt := pbdata.Header.GetWriteTime().AsTime().UTC()
	tag := pbdata.Header.Tag
	eles := make([]*sets.Element, len(pbdata.Elements), len(pbdata.Elements))
	for i, e := range pbdata.Elements {
		eles[i] = &sets.Element{
			Key: e.GetKey(),
		}
	}

	return sets.NewPrimitiveSet(wt, sets.CanonicalTag(tag), eles), nil
}

// PersistSet writes a protobuf encoded representation of the given PrimitiveSet to the provided io.Writer
func PersistSet(w io.Writer, s *sets.PrimitiveSet) error {
	now := time.Now().UTC()
	pbhead := &ReplaceStyleHeader{
		Tag:       *s.Tag(),
		WriteTime: timestamppb.New(now),
	}

	pbeles := make([]*ReplaceStyleElement, 0, s.Size())
	for e := range s.Elements() {
		pbeles = append(pbeles, &ReplaceStyleElement{
			Key: e.Key,
		})
	}

	pbfile := &ReplaceStyleElements{
		Header:   pbhead,
		Elements: pbeles,
	}

	pbwrapper := &ElementFile{
		Data: &ElementFile_Rstyle{pbfile},
	}

	unslurp, err := proto.Marshal(pbwrapper)
	if err != nil {
		return err
	}

	l, err := w.Write(unslurp)
	if err != nil {
		return err
	}

	if l < len(unslurp) {
		return errors.New("not enough bytes written while persisting")
	}

	return nil
}

func Name(ident string, pset *sets.PrimitiveSet) string {
	return fmt.Sprintf("%s_%d.pb", ident, pset.LastUpdateTime().Unix())
}
