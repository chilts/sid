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
	fmt.Printf("id1 = %s\n", sid.Id())
	fmt.Printf("id2 = %s\n", sid.Id())
	fmt.Printf("id3 = %s\n", sid.Id())
	fmt.Printf("id4 = %s\n", sid.Id())
	fmt.Printf("id5 = %s\n", sid.Id())
}
