package utlis

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func CalculateHash(timeStamp, data, previousHash string) (string, error) {
	jsonData, err := json.Marshal(fmt.Sprintf("%v%v%v", timeStamp, data, previousHash))
	if err == nil {
		return fmt.Sprintf("%X", sha256.Sum256(jsonData)), nil
	}
	return "", err
}
