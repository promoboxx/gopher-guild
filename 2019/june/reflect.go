package main

import (
	"log"
	"reflect"
	"time"
)

type Runner interface {
	Run(r interface{})
}

type runner struct {
	I int
	J int
	p int
}

func (r *runner) Run(iface interface{}) {
	kind := reflect.ValueOf(iface).Kind()

	switch kind {
	case reflect.String:
		log.Printf("iface was a string")
	case reflect.Int:
		log.Printf("iface was an int")
	}
}

func NewRunner(i, j int) Runner {
	return &runner{I: i, J: j, p: 100}
}

func main() {
	r := NewRunner(1, 2)

	// This is not possible since the compiler only sees r
	// as a Runner interface type. So we cannot access the members
	// of the private runner struct
	// log.Printf("i = %d", r.I)

	// we can assert to get the underlying type
	if value, ok := r.(*runner); ok {
		value.Run(1)
	} else {
		log.Fatalf("could not assert")
	}

	tm := time.Now()
	v := reflect.Indirect(reflect.ValueOf(tm))
	for i := 0; i < v.NumField(); i++ {
		log.Printf("value: %v", v.Field(i))
	}

	log.Printf("complete")
}
