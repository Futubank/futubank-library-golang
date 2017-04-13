package futubank

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

// Calculates a MAC (message authentication code) for requests to be sent to Futubank processing center (futubank.com)
func CalculateSignature(data map[string]string, secretKey string) string {
	// Build a string of key/value pairs
	var chunks []string
	for key, value := range data {
		if value == "" {
			continue
		}

		chunk := fmt.Sprintf("%s=%s", key, base64.StdEncoding.EncodeToString([]byte(data[key])))
		chunks = append(chunks, chunk)
	}

	// Sort chunks (by key name actually)
	sort.Strings(chunks)

	var preparedString = strings.Join(chunks, "&")

	// Calculate the MAC
	var result string
	result = sha1Hex(secretKey + preparedString)
	result = sha1Hex(secretKey + result)
	return result
}

func sha1Hex(data string) string {
	var sha1Bytes = sha1.Sum([]byte(data))
	return fmt.Sprintf("%x", sha1Bytes)
}
