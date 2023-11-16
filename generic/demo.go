package generic

import (
	"log"
	"reflect"
)

func GenericDemo[T any]() {
	var val T
	log.Println(reflect.TypeOf(val))
}

var GenericDemoBool = GenericDemo[bool]
