package dojo

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestATM(t *testing.T) {
	Convey("Withdraw 5 Euro", t, func() {
		v, e := Withdraw(5)
		So(v, ShouldResemble, map[int]int{5: 1})
		So(e, ShouldEqual, nil)
	})

	Convey("Withdraw 15 Euro", t, func() {
		v, e := Withdraw(15)
		So(v, ShouldResemble, map[int]int{5: 1, 10: 1})
		So(e, ShouldEqual, nil)
	})

	Convey("Withdraw 1595 Euro", t, func() {
		v, e := Withdraw(1595)
		So(v, ShouldResemble, map[int]int{500: 3, 50: 1, 20: 2, 5: 1})
		So(e, ShouldEqual, nil)
	})

	Convey("Withdraw 14 Euro", t, func() {
		v, e := Withdraw(14)
		So(v, ShouldResemble, map[int]int{})
		So(e, ShouldResemble, errors.New("Cannot return."))
	})

	Convey("Withdraw -5 Euro", t, func() {
		v, e := Withdraw(-5)
		So(v, ShouldResemble, map[int]int{})
		So(e, ShouldResemble, errors.New("Are you retarded?"))
	})

	Convey("Withdraw 2000 Euro", t, func() {
		v, e := Withdraw(2000)
		So(v, ShouldResemble, map[int]int{500: 4})
		So(e, ShouldResemble, nil)
	})

	Convey("Withdraw 2002 Euro", t, func() {
		v, e := Withdraw(2002)
		So(v, ShouldResemble, map[int]int{})
		So(e, ShouldResemble, errors.New("Upper limit exceeded!"))
	})
}

// func TestATM15(t *testing.T) {
// }
