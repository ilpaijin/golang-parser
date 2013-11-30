package ilpaijin

import (
	"errors"
	"log"
	"net/http"
	_ "net/http/pprof"
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
	return
}

func Parser(xmlType string, parseType string) (resultSet []reflect.Value, err error) {

	funcs := map[string]interface{}{
		"pregame": Pregame,
		"coupons": Coupons,
	}

	if funcs[xmlType] == nil {
		log.Fatal("xmlType not supported")
	}

	resultSet, err = Call(funcs, xmlType, parseType)
	if err != nil {
		log.Fatal("error calling parser interface")
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	return
}
