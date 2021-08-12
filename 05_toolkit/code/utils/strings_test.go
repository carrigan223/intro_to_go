package utils

import "testing"

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := MakeExcited("omg so exciting")
	if actual != expected {
		t.Errorf("String does not match, Actual: %s, Expected: %s", actual, expected)
	}
}