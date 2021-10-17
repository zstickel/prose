package prose

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "bannana"}, want: "apple and bannana"},
		testData{list: []string{"apple", "bannana", "orange"}, want: "apple, bannana, and orange"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Error(errorString(test.list, got, test.want))
		}
	}
}
