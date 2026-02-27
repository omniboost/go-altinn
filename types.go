package altinn

import (
	"encoding/xml"
)

type RS0297FormData struct {
	XMLName xml.Name `xml:"A3_RS0297_M"`

	DataFormatProvider string `xml:"dataFormatProvider,attr"`
	DataFormatID       string `xml:"dataFormatId,attr"`
	DataFormatVersion  string `xml:"dataFormatVersion,attr"`

	InternInfo struct {
		PeriodeNummer string `xml:"periodeNummer"`
		PeriodeAAr    string `xml:"periodeAAr"`
		EnhetsOrgNr   string `xml:"enhetsOrgNr"`
	} `xml:"InternInfo"`

	SkjemaData struct {
		SluttbrukerSystemNavn     string `xml:"sluttbrukerSystemNavn"`
		IDriftJaNei               string `xml:"iDriftJaNei"`
		AntallDagerAApen          string `xml:"antallDagerAApen"`
		TilgjengeligeGjesterom    string `xml:"tilgjengeligeGjesterom"`
		TilgjengeligeRomRullestol string `xml:"tilgjengeligeRomRullestol"`
		FasteSengeplasser         string `xml:"fasteSengeplasser"`
		Romdogn                   string `xml:"romdogn"`
		Losjiomsetning            string `xml:"losjiomsetning"`
		AntallNorskeGjester       string `xml:"antallNorskeGjester"`
		AntallUtlGjester          string `xml:"antallUtlGjester"`
		AntallGjestedogn          string `xml:"antallGjestedogn"`

		BostedsLand []struct {
			BostedsLandID        string `xml:"bostedsLandID"`
			AntallGjestedognLand string `xml:"antallGjestedognLand"`
		} `xml:"bostedsLand"`

		OppholdsFormaal []struct {
			FormaalID               string `xml:"formaalID"`
			AntallGjestedognFormaal string `xml:"antallGjestedognFormaal"`
		} `xml:"oppholdsFormaal"`
	} `xml:"SkjemaData"`
}
