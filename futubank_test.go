package futubank

import "testing"

func TestSignature(t *testing.T) {
	data := map[string]string{
		"one":   "1",
		"two":   "2",
		"empty": "",
	}

	mac := CalculateSignature(data, "secret")
	expected_mac := "07e5ae4250cff195c0afbf6740df59f40074815d"

	if mac != expected_mac {
		t.Errorf("Expected: %s, got: %s", expected_mac, mac)
	}
}
