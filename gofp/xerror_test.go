package gofp

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestXError(t *testing.T) {
	Convey("XError", t, func() {
		Convey("NewXError", func() {
			Convey("wraps actual error and returns an XError with the given message", func() {
				actualErr := fmt.Errorf("Something went wrong")
				xerr := NewXError("Something went horribly wrong", actualErr, http.StatusInternalServerError)
				So(xerr.Err, ShouldEqual, actualErr)
				So(xerr.Message, ShouldEqual, "Something went horribly wrong")
				So(xerr.ErrorCode, ShouldEqual, http.StatusInternalServerError)
			})
		})

		Convey("NewXErrorF", func() {
			Convey("returns a new XError with the formatted message", func() {
				xerr := NewXErrorF("You sent a bad request!", http.StatusBadRequest)
				So(xerr.Message, ShouldEqual, "You sent a bad request!")
				So(xerr.ErrorCode, ShouldEqual, http.StatusBadRequest)
				So(xerr.Err, ShouldNotBeNil)
			})
		})
	})
}
