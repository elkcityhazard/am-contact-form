package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
	"time"
)

type Util struct {
	hmacSecret         string
	randomStringSource string
}

func NewUtil() *Util {
	return &Util{
		hmacSecret:         "this is a really long secret phrase",
		randomStringSource: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-=+",
	}
}

//	GetCode takes in some data as a string
// 	and returns a new hmacl token

func (u *Util) CreateHmacToken(data string) (string, error) {
	h := hmac.New(sha256.New, []byte(u.hmacSecret))

	timestamp := fmt.Sprintf("%d", time.Now().Add(time.Second*3000))

	message := data + "|" + timestamp

	_, err := io.WriteString(h, message)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//	VerifyHmacToken takes in an already hashed token with a delimeter (i.e., mysecrettoken|random@domain.com)
//	delimiter is the passed in delimiter

func (u *Util) VerifyHmacToken(tokenWithDelimiter, delimiter string) bool {

	var isEqual = false

	td := strings.SplitN(tokenWithDelimiter, delimiter, 2)

	if len(td) == 2 {
		t := td[0]
		k := td[1]

		refToken, err := u.CreateHmacToken(k)

		if err != nil {
			return false
		}

		isEqual = hmac.Equal([]byte(t), []byte(refToken))

		return isEqual
	}

	return false

}

func (u *Util) GenerateRandomString(n int) string {

	randomString := make([]rune, n)

	randStringSeed := []rune(u.randomStringSource)

	for i := range randomString {
		prime, _ := rand.Prime(rand.Reader, len(randStringSeed))

		x := prime.Uint64()
		y := uint64(len(randStringSeed))

		randomString[i] = randStringSeed[x%y]

	}

	return string(randomString)

}
