package gofp

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOption(t *testing.T) {
	Convey("Option", t, func() {
		Convey("IsEmpty returns true when the option is empty", func() {
			o := None[int]()
			So(o.IsEmpty(), ShouldBeTrue)
		})

		Convey("IsEmpty returns false when the option is not empty", func() {
			o := Some(10)
			So(o.IsEmpty(), ShouldBeFalse)
		})

		Convey("HasValue returns true when the option has a value", func() {
			o := Some(10)
			So(o.HasValue(), ShouldBeTrue)
		})

		Convey("HasValue returns false when the option does not have a value", func() {
			o := None[int]()
			So(o.HasValue(), ShouldBeFalse)
		})

		Convey("Get panics when the option is empty", func() {
			o := None[int]()
			So(func() { o.Get() }, ShouldPanic)
		})

		Convey("Get returns the value when the option is not empty", func() {
			o := Some(10)
			So(o.Get(), ShouldEqual, 10)
		})

		Convey("MarshalJSON returns the JSON representation of the option", func() {
			o := Some(10)
			b, err := o.MarshalJSON()
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, "10")
		})

		Convey("MarshalJSON returns nil in JSON when the option is empty", func() {
			o := None[int]()
			b, err := o.MarshalJSON()
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, "null")
		})

		Convey("Converting None to a JSON string returns nothing", func() {
			type Payload struct {
				ID          int            `json:"id"`
				Name        string         `json:"name"`
				Description Option[string] `json:"description"`
			}
			payload := Payload{
				ID:          1,
				Name:        "test",
				Description: None[string](),
			}
			b, err := json.Marshal(payload)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, `{"id":1,"name":"test","description":null}`)
		})

		Convey("FromNilable returns an empty option when the value is nil", func() {
			var a *int
			o := FromNilable(a)
			So(o.IsEmpty(), ShouldBeTrue)
		})

		Convey("FromNilable returns an option with the value when the value is not nil", func() {
			a := new(int)
			*a = 10
			o := FromNilable(a)
			So(o.Get(), ShouldEqual, 10)
		})
	})
}
