package random

import "testing"

func TestGetStringBytes(t *testing.T) {
	strLen := 32
	randomStr := GetStringBytes(strLen)

	if strLen != len(randomStr) {
		t.Fatalf(`Expected string length: %v. Actual: %v`, strLen, len(randomStr))
	}
}
