package prose

import "testing"

func TestTwoElements(t *testing.T) {

}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	if JoinWithCommas(list) != "apple, orange, and pear" {
		t.Error("didn't match expected value")
	}
}
