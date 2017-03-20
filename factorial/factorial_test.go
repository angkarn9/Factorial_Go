package factorial_test

import (
	"testing"

	"../factorial"
)

func TestCallFactorial(t *testing.T) {
	result := factorial.Factorial(5)
	if result != 120 {
		t.Error("Expected 120, got", result)
	}
}
