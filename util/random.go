package util

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
}

// RandomProductName generates random product name
func RandomProductName() string {
	return RandomString(10)
}

// RandomBigInt generates a random integer between min and max
func RandomBigInt(min, max int64) *big.Int {
	result := big.NewInt(0)
	return result.Add(big.NewInt(min), big.NewInt(rand.Int63n(max-min+1)))
}

// RandomInt64 generates a random integer between min and max
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates random name
func RandomName() string {
	return RandomString(6)
}

// RandomSurname generates random surname
func RandomSurname() string {
	return RandomString(6)
}

// RandomAddress generates random address
func RandomStreet() string {
	return RandomString(10)
}

// RandomHomeNo generates random home number
func RandomHomeNo() string {
	return strconv.FormatInt(RandomInt64(1,100),10)

}

// RandomPLZ generates random PLZ
func RandomPLZ() string {
	return strconv.FormatInt(RandomInt64(3000,6999),10)

}

// RandomAddress generates random address
func RandomCity() string {
	return RandomString(8)
}

// RandomPhoneNumber generates random phone number
func RandomPhoneNumber() string {
	return strconv.FormatInt(RandomInt64(100000000, 999999999), 10)
}

// RandomEmail generates random email
func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(5), RandomString(4), RandomString(2))
}