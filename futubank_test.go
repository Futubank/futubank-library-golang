package futubank

import "testing"

func TestSignature(t *testing.T) {
	data := map[string]string{
		"one":   "1",
		"two":   "2",
		"empty": "",
	}

	mac := CalculateSignature(data, "secret")
	expectedMac := "07e5ae4250cff195c0afbf6740df59f40074815d"

	if mac != expectedMac {
		t.Errorf("Expected: %s, got: %s", expectedMac, mac)
	}
}
