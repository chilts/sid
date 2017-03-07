package sid

import (
	"log"
	"testing"
)

func TestId(t *testing.T) {
	id := Id()

	if len(id) != 23 {
		log.Fatal("id length should be 23")
	}

	if id[11] != '-' {
		log.Fatal("id should have a dash [-] in the middle")
	}
}
