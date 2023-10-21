package random

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomNewString(stringLength int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]rune, stringLength)

	for i := range res {
		res[i] = letters[rnd.Intn(len(letters))]
	}

	return string(res)
}
