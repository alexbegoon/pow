package proofofwork

import (
	"github.com/stretchr/testify/assert"
	"pow/internal/challenge"
	"testing"
)

func TestProofOfWork_Validate_Without_Run(t *testing.T) {
	t.Parallel()

	c := challenge.New()
	pow := NewProofOfWork(c)

	assert.False(t, pow.Validate())
}

func TestProofOfWork_Validate_With_Run(t *testing.T) {
	t.Parallel()

	c := challenge.New()
	pow := NewProofOfWork(c)

	nonce, _ := pow.Run()

	pow.Nonce = nonce

	assert.True(t, pow.Validate())
}
