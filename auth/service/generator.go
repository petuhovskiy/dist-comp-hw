package service

import (
	"crypto/rand"
	"encoding/base64"
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
	b := make([]byte, g.length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}