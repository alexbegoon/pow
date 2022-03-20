package challenge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	c := New()
	assert.NotEmpty(t, c)
}
