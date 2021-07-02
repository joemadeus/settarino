package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestData_LessThan(t *testing.T) {
	t.Run("less", func(t *testing.T) {
		left := Element{Key: "fozzie"}
		right := Element{Key: "animal"}
		assert.False(t, left.LessThan(&right))
	})

	t.Run("equal", func(t *testing.T) {
		left := Element{Key: "fozzie"}
		right := Element{Key: "fozzie"}
		assert.False(t, left.LessThan(&right))
	})

	t.Run("greater", func(t *testing.T) {
		left := Element{Key: "fozzie"}
		right := Element{Key: "kermit"}
		assert.True(t, left.LessThan(&right))
	})

	t.Run("zero_left", func(t *testing.T) {
		left := Element{Key: "animal"}
		right := Element{}
		assert.False(t, left.LessThan(&right))
	})

	t.Run("zero_right", func(t *testing.T) {
		left := Element{}
		right := Element{Key: "piggy"}
		assert.False(t, left.LessThan(&right))
	})

	t.Run("zeros_everywhere", func(t *testing.T) {
		left := Element{}
		right := Element{}
		assert.False(t, left.LessThan(&right))
	})
}

func TestData_Equals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		left := Element{Key: "janis"}
		right := Element{Key: "janis"}
		assert.True(t, left.Equals(&right))
	})

	t.Run("notequal", func(t *testing.T) {
		left := Element{Key: "rowlf"}
		right := Element{Key: "janis"}
		assert.False(t, left.Equals(&right))
	})
}
