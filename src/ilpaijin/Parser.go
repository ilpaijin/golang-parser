package ilpaijin

import (
	// "errors"
	"ilpaijin/parsers"
	"log"
	// "net/http"
	// _ "net/http/pprof"
	// "reflect"
)

// func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
// 	f := reflect.ValueOf(m[name])
// 	if len(params) != f.Type().NumIn() {
// 		err = errors.New("The number of params is not adapted.")
// 		return
// 	}
// 	in := make([]reflect.Value, len(params))
// 	for k, param := range params {
// 		in[k] = reflect.ValueOf(param)
// 	}
// 	result = f.Call(in)
// 	return
// }

func Parser(xmlType string, parseType string) (resultSet parsers.Feed) {

	var f parsers.Feed

	switch xmlType {
	case "pregame":
		resultSet = f.Pregame(parseType)
	case "coupon":
		resultSet = f.Coupons(parseType)
	default:
		log.Fatal("xmlType not supported")
	}
	return resultSet

	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
}
