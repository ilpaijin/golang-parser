package ilpaijin

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type CouponsXml struct {
	Quote []Quota `xml:Quota`
}

type Quota struct {
	TipoQuota string
	IDQuota   int
}

func (q Quota) String() string {
	return fmt.Sprintf("%s - %d", q.TipoQuota, q.IDQuota)
}

func Coupons(parseType string) CouponsXml {

	switch parseType {
	case "full":
		// run fnction full
	case "light":
		// run fnction light
	}

	xmlResource, err := os.Open("data/file_0.xml")
	if err != nil {
		log.Fatal("Error opening ", err)
	}
	defer xmlResource.Close()

	b, _ := ioutil.ReadAll(xmlResource)

	var xmlData CouponsXml
	xml.Unmarshal(b, &xmlData)

	for _, quota := range xmlData.Quote {
		fmt.Printf("\t%s\n", quota)
	}

	// log.Fatal(xmlData)

	return xmlData
}
