package gofp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFunctions(t *testing.T) {
	Convey("Functions", t, func() {
		Convey("IsNil", func() {
			Convey("returns true when the value is nil", func() {
				var a *int
				So(IsNil(a), ShouldBeTrue)
			})

			Convey("returns false when the value is not nil", func() {
				a := new(int)
				So(IsNil(a), ShouldBeFalse)
			})
		})

		Convey("IsNonNil", func() {
			Convey("returns false when the value is nil", func() {
				var a *int
				So(IsNonNil(a), ShouldBeFalse)
			})

			Convey("returns true when the value is not nil", func() {
				a := new(int)
				So(IsNonNil(a), ShouldBeTrue)
			})
		})

		Convey("Filter", func() {
			Convey("returns a new slice containing only the elements that satisfy the predicate", func() {
				source := []int{1, 2, 3, 4, 5}
				predicate := func(i int) bool { return i%2 == 0 }
				result := Filter(source, predicate)
				So(result, ShouldResemble, []int{2, 4})
			})

			Convey("returns an empty slice when the given slice is empty", func() {
				source := []int{}
				predicate := func(i int) bool { return i%2 == 0 }
				result := Filter(source, predicate)
				So(result, ShouldBeEmpty)
			})

			Convey("returns an empty slice when the given slice does not contain any elements that satisfy the predicate", func() {
				source := []int{1, 3, 5}
				predicate := func(i int) bool { return i%2 == 0 }
				result := Filter(source, predicate)
				So(result, ShouldBeEmpty)
			})
		})

		Convey("Map", func() {
			Convey("returns a new slice containing the results of applying the given function to the given slice", func() {
				f := func(i int) int { return i * 2 }
				arr := []int{1, 2, 3}
				result := Map(f, arr)
				So(result, ShouldResemble, []int{2, 4, 6})
			})

			Convey("returns an empty slice when the given slice is empty", func() {
				f := func(i int) int { return i * 2 }
				arr := []int{}
				result := Map(f, arr)
				So(result, ShouldBeEmpty)
			})
		})
	})
}
