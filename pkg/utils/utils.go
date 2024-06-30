package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
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

	message := data

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

	td := strings.SplitN(tokenWithDelimiter, delimiter, 3)

	fmt.Println(len(td))
	if len(td) == 3 {
		t := td[0]
		k := td[1]
		e, err := strconv.ParseInt(td[2], 10, 64)

		if err != nil {
			fmt.Println("no int64 convert")
			return false
		}

		fmt.Println(k)

		if time.Now().Unix() > e {
			fmt.Println("time out of bound")
			return false
		}

		refToken, err := u.CreateHmacToken(k)

		if err != nil {
			fmt.Println("couldn't convert to token")
			return false
		}

		fmt.Println(t, refToken)

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
