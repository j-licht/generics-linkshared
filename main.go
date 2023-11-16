package main

import (
	"github.com/j-licht/generics-linkshared/generic"
	"github.com/godbus/dbus/v5"
)

func main() {
	// include something to be able to do linkshared
	_ = dbus.AuthAnonymous()

	generic.GenericDemo[bool]()
}
