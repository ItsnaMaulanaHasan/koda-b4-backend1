package lib

import "github.com/matthewhartstonge/argon2"

func HashPassword(str string) ([]byte, error) {
	argon := argon2.DefaultConfig()
	hashPassword, err := argon.HashEncoded([]byte(str))
	return hashPassword, err
}
