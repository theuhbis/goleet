package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	strOne := "anagram"
	strTwo := "nagaram"
	got := isAnagram(strOne, strTwo)
	want := true

	if got != want {
		t.Errorf("expected '%v' but got '%v'", want, got)
	}
}