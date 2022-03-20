package main

import (
	"context"
	"encoding/binary"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pow/internal/challenge"
	"pow/internal/proofofwork"
	"pow/internal/quotes"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

var currentPow *proofofwork.ProofOfWork

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/challenge", GetChallenge).Methods("GET")
	router.HandleFunc("/quote", GetQuote).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

func GetChallenge(w http.ResponseWriter, _ *http.Request) {
	c := challenge.New()
	currentPow = proofofwork.NewProofOfWork(c)
	log.Println("New challenge requested")
	w.WriteHeader(200)
	w.Write(currentPow.Challenge)
}

func GetQuote(w http.ResponseWriter, r *http.Request) {

	log.Println("Quote requested")
	nonce, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if len(nonce) == 0 {
		w.WriteHeader(422)
		w.Write([]byte("Invalid request"))
		log.Println("Quote request is invalid")

		return
	}

	currentPow.Nonce = int(binary.BigEndian.Uint64(nonce))
	if !currentPow.Validate() {
		w.WriteHeader(422)
		w.Write([]byte("Invalid request"))
		log.Println("Quote request is invalid")

		return
	}

	q := quotes.GetQuote()
	w.WriteHeader(200)
	w.Write([]byte("\n\n"))
	w.Write([]byte(q))
	w.Write([]byte("\n\n"))
	log.Println("Quote request succeeded")
}
