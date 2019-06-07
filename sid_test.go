package sid

import (
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

func TestIdBase32(t *testing.T) {
	id := IdBase32()

	if len(id) != 27 {
		log.Fatal("id length should be 27")
	}

	if id[13] != '-' {
		log.Fatal("id should have a dash [-] in the middle")
	}
}

func TestIdHex(t *testing.T) {
	id := IdHex()

	if len(id) != 33 {
		log.Fatal("id length should be 33")
	}

	if id[16] != '-' {
		log.Fatal("id should have a dash [-] in the middle")
	}
}

func TestId(t *testing.T) {
	id := Id()

	if len(id) != 39 {
		log.Fatal("id length should be 39")
	}

	if id[19] != '-' {
		log.Fatal("id should have a dash [-] in the middle")
	}
}
