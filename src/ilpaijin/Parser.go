package ilpaijin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	fmt.Println(result)
	return
}

func Parser(xmlType string, parseType string) {
	funcs := map[string]interface{}{
		"pregame": Pregame,
	}

	coupon, err := Call(funcs, os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coupon[0].Interface())

}
