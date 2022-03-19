package quotes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuote(t *testing.T) {
	t.Parallel()
	q := GetQuote()
	assert.NotEmpty(t, q)
}
