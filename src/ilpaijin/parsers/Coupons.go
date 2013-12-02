package parsers

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Coupon struct {
	IDCoupon                int              `xml:"IDCoupon"`
	IDUtente                int              `xml:"IDUtente"`
	IDTipoCoupon            int              `xml:"IDTipoCoupon"`
	IDStatoCoupon           int              `xml:"IDStatoCoupon"`
	IDTipoScommessa         int              `xml:"IDTipoScommessa"`
	IDEsitoCoupon           int              `xml:"IDEsitoCoupon"`
	DataInserimento         string           `xml:"DataInserimento"`
	DataChiusura            string           `xml:"DataChiusura,attr"`
	CodiceCoupon            string           `xml:"CodiceCoupon"`
	Importo                 float64          `xml:"Importo"`
	ImportoGiocato          float64          `xml:"ImportoGiocato"`
	IDUtenteInserimento     int              `xml:"IDUtenteInserimento"`
	IDTipoUtenteInserimento int16            `xml:"IDTipoUtenteInserimento"`
	TipoScommessa           string           `xml:"TipoScommessa"`
	TipoCoupon              string           `xml:"TipoCoupon"`
	StatoCoupon             string           `xml:"StatoCoupon"`
	EsitoCoupon             string           `xml:"EsitoCoupon"`
	VincitaPotenziale       float64          `xml:"VincitaPotenziale"`
	VincitaPotenzialeMinima float64          `xml:"VincitaPotenzialeMinima"`
	IDStatoRiserva          string           `xml:"IDStatoRiserva"`
	DatiAggiuntivi          []DatoAggiuntivo `xml:"DatiAggiuntivi>DatoAggiuntivo"`
	Scommesse               []Scommessa      `xml:"Scommesse>Scommessa"`
	Quote                   []Quota          `xml:"Quote>Quota"`
}

type DatoAggiuntivo struct {
	Promozione string  `xml:"Promozione, attr"`
	Valore     float64 `xml:"Valore, attr"`
}

type Scommessa struct {
	IDScommessa int `xml:"IDScommessa"`
}

type Quota struct {
	TipoQuota string `xml:"TipoQuota"`
	IDQuota   int    `xml:"IDQuota"`
}

func (c *Coupon) Coupons(parseType string) (xmlData Feed) {

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

	xml.Unmarshal(b, &xmlData)

	return xmlData
}
