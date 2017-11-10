package pangram

import (
	"testing"
)

func TestPangramChecker_Check(t *testing.T) {
	checker := PangramChecker{Alphabet: []rune{'a', 'b', 'c'}}

	data := []struct {
		Pangram  string
		Expected bool
	}{
		{
			"abc",
			true,
		},
		{
			"cab",
			true,
		},
		{
			"abcd",
			true,
		},
		{
			"ABC",
			true,
		},
		{
			"ab",
			false,
		},
		{
			"",
			false,
		},
		{
			"💪",
			false,
		},
	}

	for _, dataElement := range data {
		r := checker.Check(dataElement.Pangram)
		if r != dataElement.Expected {
			t.Errorf("Expected  PangramChecker.Check(%s) to be %t", dataElement.Pangram, dataElement.Expected)
		}
	}
}

func BenchmarkPangramChecker_Check(b *testing.B) {
	c := PangramChecker{Alphabet: []rune{'a', 'b', 'c'}}
	for i := 0; i < b.N; i++ {
		c.Check("abcd")
	}
}
