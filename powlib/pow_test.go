package powlib

import (
	"log"
	"testing"
)

func TestSolveChallenge(t *testing.T) {
	log.Println("TestSolveChallenge")
	challenge := "test"
	nonce := SolveChallenge(challenge)
	if !IsValidPoW(challenge, nonce) {
		t.Errorf("SolveChallenge produced an invalid nonce for the challenge %s", challenge)
	}
}

func TestIsValidPoW(t *testing.T) {
	log.Println("TestIsValidPoW")
	challenge := "test"
	nonce := "93721" // A valid nonce for the challenge "test"
	if !IsValidPoW(challenge, nonce) {
		t.Errorf("IsValidPoW failed to validate a valid PoW for the challenge %s and nonce %s", challenge, nonce)
	}

	invalidNonce := "54321"
	if IsValidPoW(challenge, invalidNonce) {
		t.Errorf("IsValidPoW incorrectly validated an invalid PoW for the challenge %s and nonce %s", challenge, invalidNonce)
	}
}
