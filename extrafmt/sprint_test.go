package extrafmt

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSprintnlnSpacing(t *testing.T) {
	expected := "this is the way of the world"
	actual := Sprintnln("this", "is", "the", "way", "of", "the", "world")

	assert.Equal(t, expected, actual)
}
