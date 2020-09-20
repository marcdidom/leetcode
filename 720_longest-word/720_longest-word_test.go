package problem720

import "testing"

func TestLongestWord(t *testing.T) {
	expected := "world"
	observed := longestWord([]string{"wor", "worl", "world", "w", "wo"})
	if expected != observed {
		t.Error("Expected world got ", observed)
	}
}
