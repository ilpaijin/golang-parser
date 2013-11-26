package ilpaijin

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Coupon struct {
	IDCoupon int `xml:"IDCoupon"`
}

func Pregame(parseType string) Coupon {

	xmlFile, err := os.Open("data/file_0.xml")
	if err != nil {
		log.Fatal("Error opening ", err)
	}

	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	var coupon Coupon
	xml.Unmarshal(b, &coupon)

	return coupon
}
