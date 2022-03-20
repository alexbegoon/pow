package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"pow/internal/proofofwork"
)

func main() {
	log.Println("Init client")

	log.Println("Init challenge")
	r, err := http.Get("http://127.0.0.1:8080/challenge")

	if err != nil {
		log.Fatal(err)
	}

	c, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Challenge received %x", c))
	log.Println("Solving challenge...")

	pow := proofofwork.NewProofOfWork(c)
	nonce, hash := pow.Run()

	log.Println(fmt.Sprintf("Challenge solved: %x", hash))
	log.Println("Requesting quote...")

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(nonce))

	r, err = http.Post("http://127.0.0.1:8080/quote", "", bytes.NewReader(b))

	if err != nil {
		log.Fatal(err)
	}

	res, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Received a quote: %v", string(res)))
}
