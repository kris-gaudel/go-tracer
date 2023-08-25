package interval

import "testing"

func Test(t *testing.T) {
	interval := Interval{Min: -10.0, Max: 10.0}

	// Test Contains
	resultContains := interval.Contains(10.0)
	expectedContains := true
	if resultContains != expectedContains {
		t.Errorf("interval.Contains(0.0) = %t; expected %t", resultContains, expectedContains)
	}

	// Test Surrounds
	resultSurrounds := interval.Surrounds(0.0)
	expectedSurrounds := true
	if resultSurrounds != expectedSurrounds {
		t.Errorf("interval.Surrounds(0.0) = %t; expected %t", resultSurrounds, expectedSurrounds)
	}

	resultSurroundsFalse := interval.Surrounds(10.0)
	expectedSurroundsFalse := false
	if resultSurroundsFalse != expectedSurroundsFalse {
		t.Errorf("interval.Surrounds(10.0) = %t; expected %t", resultSurroundsFalse, expectedSurroundsFalse)
	}

	// Test Clamp
	resultClamp := interval.Clamp(11.0)
	expectedClamp := 10.0
	if resultClamp != expectedClamp {
		t.Errorf("interval.Clamp(0.0) = %f; expected %f", resultClamp, expectedClamp)
	}
}
