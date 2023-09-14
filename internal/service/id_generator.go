package service

import "math/rand"

type IdGenerator interface {
	New(length int) string
}

type idGeneratorImpl struct{}

func (i idGeneratorImpl) New(length int) string {
	// 48 - 57  => 0 - 9
	// 65 - 90 => A - Z
	// 97 - 122 => a - z
	res := ""
	for i := 0; i < length; i++ {
		j := rand.Intn(3)
		if j == 0 {
			// 0-9
			res += string(byte(48 + rand.Intn(10)))
		}
		if j == 1 {
			// A-Z
			res += string(byte(65 + rand.Intn(26)))
		}

		if j == 2 {
			// a-z
			res += string(byte(97 + rand.Intn(26)))
		}
	}
	return res
}

func NewIdGenerator() IdGenerator {
	return &idGeneratorImpl{}
}
