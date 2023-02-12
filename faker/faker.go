package faker

import (
	"fmt"
	"math/rand"
	"time"
	"twitter/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var Password = "WAgD4aXm6kJhgszKJYMhW84b"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RndStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes)/2)]
	}

	return string(b)
}

func RandInt(min, max int) int {
	return rand.Intn(max - min + 1)
}

func Username() string {
	return fmt.Sprintf("%s", RndStringRunes(RandInt(3, 10)))
}

func ID() string {
	return fmt.Sprintf("%s-%s-%s-%s", RndStringRunes(4), RndStringRunes(4), RndStringRunes(4), RndStringRunes(4))
}

func UUID() string {
	return uuid.Generate()
}

func Email() string {
	return fmt.Sprintf("%s@gmail.com", RndStringRunes(RandInt(5, 10)))
}

func RandStr(n int) string {
	return RndStringRunes(n)
}
