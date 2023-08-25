package utils

import "testing"

func Test(t *testing.T) {
	// Test DegreesToRadians
	result := DegreesToRadians(180.0)
	expected := PI

	if result != expected {
		t.Errorf("DegreesToRadians(180.0) = %f; expected %f", result, expected)
	}

	// Test RandomDouble
	result = RandomDouble()
	expected = RandomDouble()
	if result == expected {
		t.Errorf("RandomDouble() = %f; expected %f (both are the same)", result, expected)
	}

	// Test RandomDoubleRange
	result = RandomDoubleRange(0.0, 10.0)
	expected = RandomDoubleRange(0.0, 10.0)
	if result == expected {
		t.Errorf("RandomDoubleRange(0.0, 10.0) = %f; expected %f (both are the same)", result, expected)
	}
}
