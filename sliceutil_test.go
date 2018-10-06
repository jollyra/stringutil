package stringutil

import "testing"

func TestIndexOf(t *testing.T) {
	cases := []struct {
		in0  []string
		in1  string
		want int
	}{
		{
			[]string{"a", "b", "c"},
			"c",
			2,
		}, {
			[]string{"a", "b", "b"},
			"b",
			1,
		}, {
			[]string{"a", "b", "c"},
			"d",
			-1,
		}, {
			[]string{},
			"d",
			-1,
		},
	}
	for _, c := range cases {
		got := IndexOf(c.in0, c.in1)
		if got != c.want {
			t.Errorf("IndexOf(%s, %s) == %d, want %d",
				c.in0, c.in1, got, c.want)
		}
	}
}
