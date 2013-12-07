package ilpaijin

import (
	"ilpaijin/parsers"
	"log"
)

func Parse(xmlType string, parseType string) (resultSet parsers.Feed) {

	switch xmlType {
	case "pregame":
		f := new(parsers.ArrayOfSport)
		resultSet = f
	case "coupon":
		f := new(parsers.Coupon)
		resultSet = f
	default:
		log.Fatal("xmlType not supported")
	}

	resultSet.Parse(parseType)

	return
}
