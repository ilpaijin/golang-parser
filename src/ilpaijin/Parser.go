package ilpaijin

import (
	"ilpaijin/parsers"
	"log"
)

func Parse(xmlType string, parseType string) (resultSet parsers.Feed) {

	switch xmlType {
	case "pregame":
		f := new(parsers.ArrayOfSport)
		resultSet.Data = f.Parse(parseType)
	case "coupon":
		f := new(parsers.Coupon)
		resultSet.Data = f.Parse(parseType)
	default:
		log.Fatal("xmlType not supported")
	}

	log.Fatal(resultSet.Data)

	return
}
