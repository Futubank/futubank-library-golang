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
	// Create a sorted list of keys
	keys := make([]string, len(data))
	i := 0
	for k, _ := range data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// Build a string of key/value pairs
	var chunks []string
	for _, key := range keys {
		if data[key] == "" {
			continue
		}

		chunk := fmt.Sprintf("%s=%s", key, base64.StdEncoding.EncodeToString([]byte(data[key])))
		chunks = append(chunks, chunk)
	}
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
