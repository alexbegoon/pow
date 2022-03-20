package challenge

import (
	"math/rand"
	"time"
)

type Challenge []byte

func New() Challenge {
	rand.Seed(time.Now().UnixMicro())

	challenge := make([]byte, 8)
	rand.Read(challenge)

	return challenge
}
