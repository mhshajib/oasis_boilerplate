package utils

import (
	"math/rand"
	"time"
)

// RandomFromStrSlice return random from a string slice
func RandomFromStrSlice(ss []string) string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(ss)
	return ss[n]
}
