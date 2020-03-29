package service

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var (
	lowercaseChars = "qwertyuiopasdfghjklzxcvbnm"
	allChars = []rune(lowercaseChars + strings.ToUpper(lowercaseChars))
)

type Generator struct {
	length uint
}

func NewGenerator(length uint) *Generator {
	return &Generator{
		length: length,
	}
}

func (g *Generator) Generate() (string, error) {
	res := make([]rune, g.length)

	max := big.NewInt(int64(len(allChars)))

	for i := range res {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		res[i] = allChars[num.Int64()]
	}

	return string(res), nil
}