package main

import "testing"

func TestSignature(t *testing.T) {
	CalculateSignature(map[string]string{
		"one":   "two",
		"three": "four",
	}, "abcd")
}
