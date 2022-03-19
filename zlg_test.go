package zlg_test

import (
	"errors"

	"github.com/kpurdon/zlg"
)

func ExampleLogger_Info() {
	zlg.New().With("foo", "bar").Info("hello")
}

func ExampleLogger_Error() {
	zlg.New().With("foo", "bar").Error(errors.New("test"))
}
