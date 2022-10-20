package twosum

import "testing"

func Equal(got, want []int) bool {
    if len(got) != len(want) {
        return false
    }
    for i, v := range got {
        if v != want[i] {
            return false
        }
    }
    return true
}

func TestTwoSum(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	got := TwoSum(nums, target)
	want := []int{0,1}

	if ! Equal(got, want) {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}