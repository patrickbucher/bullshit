package bullshit

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(strings []string) []string {
	backup := make([]string, len(strings))
	copy(backup, strings)
	shuffled := make([]string, len(backup))
	indices := r.Perm(len(backup))
	var s int
	for _, i := range indices {
		shuffled[i] = backup[s]
		s++
	}
	return shuffled
}
