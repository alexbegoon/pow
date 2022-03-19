package proofofwork

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProofOfWork_Validate_Without_Run(t *testing.T) {
	t.Parallel()

	pow := NewProofOfWork()

	assert.False(t, pow.Validate())
}

func TestProofOfWork_Validate_With_Run(t *testing.T) {
	t.Parallel()

	pow := NewProofOfWork()

	nonce, _ := pow.Run()

	pow.Nonce = nonce

	assert.True(t, pow.Validate())
}
