package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateUniqueOTP generates a unique OTP consisting of a 3-letter prefix and a 6-digit number.
func GenerateUniqueOTP() (string, int, error) {
	prefix, err := generateRandomPrefix()
	if err != nil {
		return "", 0, err
	}

	otp, err := generateRandomNumber(100000, 999999)
	if err != nil {
		return "", 0, err
	}

	return prefix, otp, nil
}

// generateRandomPrefix generates a random 3-letter uppercase prefix.
func generateRandomPrefix() (string, error) {
	var prefix string
	for i := 0; i < 3; i++ {
		num, err := generateRandomNumber(65, 90) // ASCII values for 'A' and 'Z'
		if err != nil {
			return "", err
		}
		prefix += string(rune(num))
	}
	return prefix, nil
}

// generateRandomNumber generates a random integer between min and max.
func generateRandomNumber(min int, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}
