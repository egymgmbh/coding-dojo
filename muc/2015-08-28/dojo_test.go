package dojo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProblem(t *testing.T) {
	Convey("Offensive word expected.", t, func() {
		So(Problem("snond", "snond"), ShouldBeTrue)
	})

	Convey("No offensive word expected.", t, func() {
		So(Problem("mispronounced", "snond"), ShouldBeFalse)
	})

	Convey("Offensive word expected.", t, func() {
		So(Problem("synchronized", "snond"), ShouldBeTrue)
	})

	Convey("Offensive word expected.", t, func() {
		So(Problem("misfunctioned", "snond"), ShouldBeTrue)
	})

	Convey("No offensive word expected.", t, func() {
		So(Problem("shotgunned", "snond"), ShouldBeFalse)
	})

	Convey("No offensive word expected.", t, func() {
		So(Problem("mispronounced", "snond"), ShouldBeFalse)
	})

	Convey("Offensive word expected.", t, func() {
		So(Problem("dnons", "snond"), ShouldBeFalse)
	})

	Convey("Offensive word expected.", t, func() {
		So(Problem("synchronized", "dnnos"), ShouldBeFalse)
	})
}
