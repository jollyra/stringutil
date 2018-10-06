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

func TestEqualSlice(t *testing.T) {
	cases := []struct {
		in0, in1 []string
		want     bool
	}{
		{
			[]string{"a", "b"},
			[]string{"a", "b"},
			true,
		}, {
			[]string{"a", "c"},
			[]string{"a", "b"},
			false,
		}, {
			[]string{"a", "b", "c"},
			[]string{"a", "b"},
			false,
		}, {
			[]string{"b", "a"},
			[]string{"a", "b"},
			true,
		}, {
			[]string{},
			[]string{},
			true,
		},
	}

	for _, c := range cases {
		got := EqualSlice(c.in0, c.in1)
		if got != c.want {
			t.Errorf("EqualSlice(%v, %v) == %v, want %v",
				c.in0, c.in1, got, c.want)
		}
	}
}

func TestUnique(t *testing.T) {
	cases := []struct {
		in   []string
		want []string
	}{
		{
			[]string{"a", "a", "b", "b"},
			[]string{"a", "b"},
		}, {
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
	}

	for _, c := range cases {
		got := Unique(c.in)
		if !EqualSlice(got, c.want) {
			t.Errorf("Unique(%v) == %v, want %v",
				c.in, got, c.want)
		}
	}
}
