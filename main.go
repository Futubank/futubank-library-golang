package main

import "sort"
import "bytes"
import "encoding/base64"
import "crypto/sha1"
import "fmt"

func CalculateSignature(data map[string]string, secret_key string) string {
	keys := make([]string, len(data))
	i := 0
	for k, _ := range data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	var buffer bytes.Buffer

	for i, key := range keys {
		if i > 0 {
			buffer.WriteString("&")
		}
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(data[key])))
	}
	fmt.Println(buffer.String())

	var result []byte

	result = append(result, []byte(secret_key)...)
	result = append(result, buffer.Bytes()...)
	result = sha1.Sum(result)
	return string(result)
}
