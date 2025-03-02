package gofp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEither(t *testing.T) {
	Convey("Either", t, func() {
		Convey("IsRight returns false and getting Right panics when the left object is not nil", func() {
			e := Left[int](NewXErrorF("some error", 500))
			So(e.IsRight(), ShouldBeFalse)
			So(e.IsLeft(), ShouldBeTrue)
			So(e.Left(), ShouldNotBeNil)
		})

		Convey("getting Right panics when the left object is not nil", func() {
			e := Left[int](NewXErrorF("some error", 500))
			So(func() { e.Right() }, ShouldPanic)
		})

		Convey("returns correct value when Right contains a value", func() {
			e := Right(10)
			So(e.IsRight(), ShouldBeTrue)
			So(e.IsLeft(), ShouldBeFalse)
			So(e.Right(), ShouldEqual, 10)
			So(func() { e.Left() }, ShouldPanic)
		})
	})
}
