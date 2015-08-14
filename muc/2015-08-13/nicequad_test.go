package nicequad

import "testing"

func TestPointQuadrant(t *testing.T) {
	if (PointQuadrant(point{1, 2}) != 0) {
		t.Errorf("Values of beta will give rise to dom! first")
	}
	if PointQuadrant(point{-1, 2}) != 1 {
		t.Errorf("Values of beta will give rise to dom! second")
	}

	if PointQuadrant(point{-2, -1}) != 2 {
		t.Errorf("Values of beta will give rise to dom! 3rd")
	}

	if PointQuadrant(point{1, -2}) != 3 {
		t.Errorf("Values of beta will give rise to dom! 4th")
	}

}

func TestIsNicequad(t *testing.T) {
	q1p1 := point{1, 1}
	q1p2 := point{-1, -1}
	q1p3 := point{1, -1}
	q1p4 := point{-1, 1}

	q2p1 := point{2, 1}
	q2p2 := point{-2, -1}
	q2p3 := point{1, -1}
	q2p4 := point{-1, 1}

	q3p1 := point{2, 1}
	q3p2 := point{2, -1}
	q3p3 := point{1, -1}
	q3p4 := point{43, 21}

	if !IsNicequad([4]point{q1p1, q1p2, q1p3, q1p4}) {
		t.Errorf("This was actually a nice quadrand")
	}

	if !IsNicequad([4]point{q2p1, q2p2, q2p3, q2p4}) {
		t.Errorf("This was actually a nice quadrand")
	}

	if IsNicequad([4]point{q3p1, q3p2, q3p3, q3p4}) {
		t.Errorf("This was actually NOT a nice quadrand")
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, AB", "BA ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
