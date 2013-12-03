package ilpaijin

import (
	"ilpaijin/parsers"
	"log"
)

func Parse(xmlType string, parseType string) (resultSet parsers.Feed) {

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
}
