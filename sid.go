// Package sid provides the ability to generate Sortable Identifiers. These are also universally unique.
//
//     id1 := sid.Id()
//     id2 := sid.Id()
//     fmt.Printf("id1 = %s\n", id1)
//     fmt.Printf("id2 = %s\n", id2)
//     // -> "id1 = 1IeSBAWW83j-2wgJ4PUtlAr"
//     // -> "id2 = 1IeSBAWW9kK-0cDG64GQgGJ"
//
// This package is simple and only provides one function. The aim here is not pure speed, it is for an easy use-case
// without having to worry about goroutines and locking.
package sid

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Remember the lastTime so that if (by chance) we get the same NanoSecond, we just incrememt the last random number.
var lastTime int64
var lastRand int64
var chars = make([]string, 11, 11)
var mu = &sync.Mutex{}

// 64 chars but ordered by ASCII
const base64 string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

func toStr(now int64) string {
	// now do the generation (backwards, so we just %64 then /64 along the way)
	for i := 10; i >= 0; i-- {
		index := now % 64
		chars[i] = string(base64[index])
		now = now / 64
	}

	return strings.Join(chars, "")
}

// Id returns a 23 char string based on timestamp and a random number. The format is "XXXXXXXXXXX-YYYYYYYYYYY" where X
// is the timestamp and Y is the random number. If (by any chance) this is called in the same nanosecond, the random
// number is incremented instead of a new one being generated. This makes sure that two consecutive Ids generated in
// the same goroutine also ensure those Ids are also sortable.
//
// It is safe to call from different goroutines since it has it's own locking.
func Id() string {
	// lock for lastTime, lastRand, and chars
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().UTC().UnixNano()
	var r int64

	// if we have the same time, just inc lastRand, else create a new one
	if now == lastTime {
		lastRand++
	} else {
		lastRand = rand.Int63()
	}
	r = lastRand

	// remember this for next time
	lastTime = now

	return toStr(now) + "-" + toStr(r)
}
