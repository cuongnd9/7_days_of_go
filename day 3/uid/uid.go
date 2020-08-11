package uid

import (
	"math/rand"
	"time"
)

/**
Generate unique id
*/
func Uid(length int) string {
	rand.Seed(time.Now().UnixNano())
	if length <= 0 {
		length = 16
	}
	return randomString(length)
}

/**
Create a random string with length
*/
func randomString(length int) string {
	bytes := make([]byte, length)
	pool := "-_@#$%^&*"
	for i := 0; i < length; i++ {
		if rand.Intn(len(pool)) == 0 {
			bytes[i] = pool[rand.Intn(len(pool))]
		} else {
			bytes[i] = byte(randomInt(97, 122))
		}
	}
	return string(bytes)
}

/**
Create a random integer from min to max - 1
*/
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

