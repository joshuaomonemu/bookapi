package utils

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
)

func Jparse(str struct{}) ([]byte, error) {
	bs, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func GenerateBookingNumber(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}
