package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCalculator(t *testing.T) {
	Convey("1 evaluates to 1", t, func() {
		So(Calculate("1"), ShouldEqual, 1)
	})

	Convey("1 + 1 evaluates to 2", t, func() {
		So(Calculate("1 + 1"), ShouldEqual, 2)
	})

	Convey("1 + 1 evaluates to 2", t, func() {
		So(Calculate("1 + 6 + 2"), ShouldEqual, 9)
	})

	Convey("1 + 1 evaluates to 2", t, func() {
		So(Calculate("1 + 42 - 1"), ShouldEqual, 42)
	})

	Convey("1 tokenizes to ['1']", t, func() {
		So(Tokenize("1"), ShouldResemble, []string{"1"})
	})

	Convey("1 + 1 tokenizes to ['1', '+', '1']", t, func() {
		So(Tokenize("1+1"), ShouldResemble, []string{"1", "+", "1"})
	})

	Convey("1 + 1 tokenizes to ['1', '+', '1']", t, func() {
		So(Tokenize("1 + 1"), ShouldResemble, []string{"1", "+", "1"})
	})

	Convey("2 + 3 tokenizes to ['2', '+', '3']", t, func() {
		So(Tokenize("2 + 3"), ShouldResemble, []string{"2", "+", "3"})
	})

	Convey("2 * 3 tokenizes to ['2', '*', '3']", t, func() {
		So(Tokenize("2 * 3"), ShouldResemble, []string{"2", "*", "3"})
	})

	Convey("2 / 3 tokenizes to ['2', '/', '3']", t, func() {
		So(Tokenize("2 / 3"), ShouldResemble, []string{"2", "/", "3"})
	})

	Convey("(2 / 3) + 1 tokenizes to ['(', '2', '/', '3', ')', '+', '1']", t, func() {
		So(Tokenize("(2 / 3) + 1"), ShouldResemble, []string{"(", "2", "/", "3", ")", "+", "1"})
	})

	Convey("(2 / 3) + 1 + (2*3) tokenizes to ['(', '2', '/', '3', ')', '+', '1', '*', '(', '2', '*', '2', '3']", t, func() {
		So(Tokenize("(2 / 3) + 1 + (2*3) "), ShouldResemble, []string{"(", "2", "/", "3", ")", "+", "1", "+", "(", "2", "*", "3", ")"})
	})

	Convey("transform 1 + 1", t, func() {
		a := []int{1, 1}
		b := []string{"+"}

		r1, r2 := Transform(Tokenize("1 + 1"))
		So(r1, ShouldResemble, b)
		So(r2, ShouldResemble, a)
	})

	Convey("transform 1 + 1 - 1", t, func() {
		a := []int{1, 1, 1}
		b := []string{"+", "-"}
		r1, r2 := Transform(Tokenize("1 + 1 - 1"))
		So(r1, ShouldResemble, b)
		So(r2, ShouldResemble, a)
	})
}
