// Package stringutil contains utility functions for working with strings.
package nicequad

type point struct {
	x int
	y int
}

func PointQuadrant(p point) int {
	if p.x >= 0 && p.y >= 0 {
		return 0
	} else if p.x >= 0 && p.y < 0 {
		return 3
	} else if p.x < 0 && p.y < 0 {
		return 2
	}

	return 1
}

func IsNicequad(points [4]point) bool {
	quadrants := [4]bool{false, false, false, false}
	for _, p := range points {
		quadrants[PointQuadrant(p)] = true
	}

	for _, q := range quadrants {
		if !q {
			return false
		}
	}

	return true
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
