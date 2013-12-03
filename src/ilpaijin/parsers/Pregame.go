package parsers

import (
	"encoding/csv"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type ArrayOfSport struct {
	XMLName xml.Name `xml:"ArrayOfSport"`
	Sports  []Sports `xml:"Sport"`
}

type Sports struct {
	IDSport    int16    `xml:"IDSport,attr"`
	Sport      string   `xml:"Sport,attr"`
	GruppiList []Gruppi `xml:"Gruppi>Gruppi"`
}

type Gruppi struct {
	IDGruppo   int16    `xml:"IDGruppo,attr"`
	Gruppo     string   `xml:"Gruppo,attr"`
	EventiList []Eventi `xml:"Eventi>Eventi"`
}

type Eventi struct {
	IDEvento        int           `xml:"IDEvento,attr"`
	Evento          string        `xml:"Evento,attr"`
	IDTipoEvento    int16         `xml:"IDTipoEvento"`
	SottoEventiList []SottoEventi `xml:"SottoEventi>SottoEventi"`
}

type SottoEventi struct {
	IDSottoEvento    int     `xml:"IDSottoEvento,attr"`
	CodPubblicazione int16   `xml:"CodPubblicazione,attr"`
	SottoEvento      string  `xml:"SottoEvento,attr"`
	DataInizio       string  `xml:"DataInizio,attr"`
	Quotelist        []Quote `xml:"Quote>Quote"`
}

type Quote struct {
	IDQuota     int     `xml:"IDQuota,attr"`
	IDTipoQuota int16   `xml:"IDTipoQuota,attr"`
	Quota       float32 `xml:"Quota,attr"`
	Giocabilita int8    `xml:"Giocabilita,attr"`
	HND         float32 `xml:"DataInizio,attr"`
	TipoQuota   string  `xml:"TipoQuota,attr"`
	ClasseQuota string  `xml:"ClasseQuota,attr"`
}

func (p *ArrayOfSport) Pregame(parseType string) (xmlData Feed) {

	xmlResource, err := os.Open("data/" + parseType + "_pregame.xml")
	if err != nil {
		log.Fatal("Error opening ", err)
	}

	switch parseType {
	case "full":
		// run fnction full
	case "light":
		// run fnction light
	}

	defer xmlResource.Close()

	b, _ := ioutil.ReadAll(xmlResource)

	xml.Unmarshal(b, &xmlData)

	return xmlData
}

func (c *ArrayOfSport) Csvize() {
	file, err := os.Create("data/Pregame.csv")
	if err != nil {
		log.Fatal("Error creating csv file")
	}
	defer file.Close()

	log.Fatal(c)

	writer := csv.NewWriter(file)

	var record []string

	for _, it := range c.Sports {

		record = append(record, string(it.IDSport))

		writer.Write(record)
	}

	writer.Flush()
}
