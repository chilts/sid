package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chilts/sid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())
	fmt.Printf("id = %s\n", sid.Id())

	fmt.Println()

	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())
	fmt.Printf("b64 = %s\n", sid.IdBase64())

	fmt.Println()

	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())
	fmt.Printf("b32 = %s\n", sid.IdBase32())

	fmt.Println()

	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
	fmt.Printf("hex = %s\n", sid.IdHex())
}
