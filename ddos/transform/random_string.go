package transform

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func ReplaceWithRandomString(str string) string {
	reg := regexp.MustCompile("randomString\\((\\d+),(\\d+)\\)")
	submatch := reg.FindStringSubmatch(str)

	from, err := strconv.Atoi(submatch[1])
	if err != nil {
		panic(err)
	}

	to, err := strconv.Atoi(submatch[2])
	if err != nil {
		panic(err)
	}

	if from > to {
		panic("Invalid randomString args")
	}

	return GenerateRandomString(from, to)
}

func GenerateRandomString(from int, to int) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var length = random.Intn(to-from) + from
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}
