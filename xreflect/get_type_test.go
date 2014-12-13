package xreflect

import (
	"testing"

	"github.com/bmizerany/assert"
	log "github.com/stuartherbert/go_modlog"
)

type getTypeExample1 struct{}
type getTypeExample2 uint8

func TestCanGetType(t *testing.T) {
	expecteds := []string{
		"github.com/stuartherbert/go_extras/xreflect.getTypeExample1",
		"github.com/stuartherbert/go_extras/xreflect.getTypeExample2",
		"github.com/stuartherbert/go_modlog.Logger",
	}

	inputs := []interface{}{
		getTypeExample1{},
		getTypeExample2(10),
		log.Logger{},
	}

	for i, input := range inputs {
		actual := get_type(input)
		assert.Equal(t, expecteds[i], actual)
	}
}
