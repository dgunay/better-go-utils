package conversion_test

import (
	"testing"

	"github.com/dgunay/better-go-utils/conversion"
)

type Foo struct{}

func (f Foo) Into() Foo { return f }

type ConvertibleIntoFoo struct{}

func (c ConvertibleIntoFoo) Into() Foo {
	return Foo{}
}

type NotConvertibleIntoFoo struct{}

func TakesAnythingConvertibleToFoo(foo conversion.Into[Foo]) {

}

func TestConversion(t *testing.T) {
	t.Run("takes anything convertible into Foo", func(t *testing.T) {
		TakesAnythingConvertibleToFoo(Foo{})
		TakesAnythingConvertibleToFoo(ConvertibleIntoFoo{})
		// TakesAnythingConvertibleToFoo(NotConvertibleIntoFoo{})
	})
}
