package sid

import (
	"fmt"
	"log"
	"testing"
)

func TestIdBase64(t *testing.T) {
	id := IdBase64()

	if len(id) != 23 {
		log.Fatal("id length should be 23")
	}

	if id[11] != '-' {
		log.Fatal("id should have a dash [-] in the middle")
	}
}

func TestId(t *testing.T) {
	id := Id()

	fmt.Printf("i=%d\n", len(id))

	if len(id) != 39 {
		log.Fatal("id length should be 39")
	}

	if id[19] != '-' {
		log.Fatal("id should have a dash [-] in the middle")
	}
}
