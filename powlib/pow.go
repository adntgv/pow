package powlib

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"
)

// SolveChallenge performs Proof of Work on the given challenge string and returns the nonce.
func SolveChallenge(challenge string) string {
	for i := 0; ; i++ {
		hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", challenge, i)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), "0000") {
			log.Printf("Found nonce %d for challenge %s", i, challenge)
			return fmt.Sprintf("%d", i)
		}
	}
}

// IsValidPoW verifies if the provided challenge and nonce satisfy the Proof of Work requirements.
func IsValidPoW(challenge string, nonce string) bool {
	log.Printf("Verifying nonce %s for challenge %s", nonce, challenge)
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s%s", challenge, nonce)))
	return strings.HasPrefix(fmt.Sprintf("%x", hash), "0000")
}
