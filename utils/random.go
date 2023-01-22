package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func GenrateRandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomCurrency() string {
	currencies := []string{
		CAD, EUR, GPD, USD,
	}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomInt() int64 {
	return randomInt(0, 10000)
}

func RandomOwner() string {
	return GenrateRandomString(10)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", GenrateRandomString(6))
}
