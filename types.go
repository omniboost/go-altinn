package altinn

import (
	"encoding/xml"
	"io"
	"strconv"
)

type Melding struct {
	XMLName            xml.Name `xml:"melding"`
	Seres              string   `xml:"seres,attr"`
	DataFormatProvider string   `xml:"dataFormatProvider,attr"`
	DataFormatId       string   `xml:"dataFormatId,attr"`
	DataFormatVersion  string   `xml:"dataFormatVersion,attr"`
	InternInformasjon  struct {
		PeriodeFritekst string `xml:"periodeFritekst"`
		RapportPeriode  struct {
			FomDato string `xml:"fomDato"`
			TomDato string `xml:"tomDato"`
		} `xml:"rapportPeriode"`
		RaNummer                  string `xml:"raNummer"`
		DelRegNummer              string `xml:"delRegNummer"`
		IdentnummerEnhet          string `xml:"identnummerEnhet"`
		SendtFraSluttbrukersystem string `xml:"sendtFraSluttbrukersystem"`
		StatistiskEnhet           struct {
			Enhetsident string `xml:"enhetsident"`
			Enhetstype  string `xml:"enhetstype"`
		} `xml:"statistiskEnhet"`
		Skjemaidentifikasjon struct {
			Periodenummer        string `xml:"periodenummer"`
			Periodetype          string `xml:"periodetype"`
			PeriodeAAr           string `xml:"periodeAAr"`
			Undersoekelsesnummer string `xml:"undersoekelsesnummer"`
		} `xml:"skjemaidentifikasjon"`
	} `xml:"InternInformasjon"`
	KontaktpersonOgKommentarfelt struct {
		Kontaktperson struct {
			Epostadresse string `xml:"epostadresse"`
			Navn         string `xml:"navn"`
			TelefonSFU   string `xml:"telefonSFU"`
		} `xml:"kontaktperson"`
	} `xml:"KontaktpersonOgKommentarfelt"`
	ForetakOgVirksomhetsopplysninger struct {
		Virksomhet struct {
			OrganisasjonsnummerVirksomhet string `xml:"organisasjonsnummerVirksomhet"`
			NavnVirksomhet                string `xml:"navnVirksomhet"`
			AdresseVirksomhet             struct {
				Gateadresse string `xml:"gateadresse"`
				Postnummer  string `xml:"postnummer"`
				Poststed    string `xml:"poststed"`
			} `xml:"adresseVirksomhet"`
		} `xml:"virksomhet"`
	} `xml:"ForetakOgVirksomhetsopplysninger"`
	StatusVirksomhetMedDriftsperiode struct {
		DriftsstatusPeriode string `xml:"driftsstatusPeriode"`
	} `xml:"StatusVirksomhetMedDriftsperiode"`
	Naeringskontrollspoersmaal struct {
		VisNaeringskontrollJaNeiPrefill string `xml:"visNaeringskontrollJaNeiPrefill"`
		Naeringskontroll                struct {
			Naeringskode          string `xml:"naeringskode"`
			Naeringstekst         string `xml:"naeringstekst"`
			Naeringsbeskrivelse   string `xml:"naeringsbeskrivelse"`
			NyNaeringsbeskrivelse struct {
				AlltidViktigsteAktivitet string `xml:"alltidViktigsteAktivitet"`
			} `xml:"nyNaeringsbeskrivelse"`
		} `xml:"Naeringskontroll"`
	} `xml:"Naeringskontrollspoersmaal"`
	Oppgavebyrde struct {
		VisOppgavebyrdeJaNeiPrefill    string `xml:"visOppgavebyrdeJaNeiPrefill"`
		TidsbrukSamleInformasjon       string `xml:"tidsbrukSamleInformasjon"`
		TidsbrukSelveSkjemautfyllingen string `xml:"tidsbrukSelveSkjemautfyllingen"`
		TidsbrukTotalHjelpAvAndre      string `xml:"tidsbrukTotalHjelpAvAndre"`
	} `xml:"Oppgavebyrde"`
	Skjemadata struct {
		Rapporteringsenhet struct {
			TypeHotell     string `xml:"typeHotell"`
			TypeVandrehjem string `xml:"typeVandrehjem"`
			TypeHytter     string `xml:"typeHytter"`
			TypeCamping    string `xml:"typeCamping"`
			TypeAnnen      string `xml:"typeAnnen"`
		} `xml:"rapporteringsenhet"`
		Kapasitet struct {
			KapasitetHotell struct {
				AntDager struct {
					Endret       string `xml:"endret"`
					AntallEndret string `xml:"antallEndret"`
				} `xml:"antDager"`
				AntGjesterom struct {
					Endret       string `xml:"endret"`
					AntallEndret string `xml:"antallEndret"`
				} `xml:"antGjesterom"`
				AntRullestol struct {
					Endret       string `xml:"endret"`
					AntallEndret string `xml:"antallEndret"`
				} `xml:"antRullestol"`
				AntFasteSengeplasser struct {
					Endret       string `xml:"endret"`
					AntallEndret string `xml:"antallEndret"`
				} `xml:"antFasteSengeplasser"`
			} `xml:"kapasitetHotell"`
			KapasitetHytte struct {
				AntDager             string `xml:"antDager"`
				AntUtleieenheter     string `xml:"antUtleieenheter"`
				AntFasteSengeplasser string `xml:"antFasteSengeplasser"`
			} `xml:"kapasitetHytte"`
			KapasitetCamping struct {
				AntDager             string `xml:"antDager"`
				AntUtekapasitet      string `xml:"antUtekapasitet"`
				AntInnekapasitet     string `xml:"antInnekapasitet"`
				AntFasteSengeplasser string `xml:"antFasteSengeplasser"`
			} `xml:"kapasitetCamping"`
		} `xml:"kapasitet"`
		OmsetningHotell struct {
			Romdoegn       string `xml:"romdoegn"`
			Losjiomsetning string `xml:"losjiomsetning"`
		} `xml:"omsetningHotell"`
		Belegg struct {
			BeleggHotell struct {
				AntGjester             string `xml:"antGjester"`
				AntNorskeGjester       string `xml:"antNorskeGjester"`
				AntGjestedoegn         string `xml:"antGjestedoegn"`
				FerieFritidGjestedoegn string `xml:"ferieFritidGjestedoegn"`
				KursGjestedoegn        string `xml:"kursGjestedoegn"`
				ForretningsGjestedoegn string `xml:"forretningsGjestedoegn"`
				Land                   struct {
					Norden  IntList `xml:"norden"`
					Europa  IntList `xml:"europa"`
					Asia    IntList `xml:"asia"`
					Amerika IntList `xml:"amerika"`
					Oceania IntList `xml:"oceania"`
					Afrika  IntList `xml:"afrika"`
				} `xml:"land"`
			} `xml:"beleggHotell"`
		} `xml:"belegg"`
	} `xml:"Skjemadata"`
}

type IntList map[string]int

func (c *IntList) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	*c = IntList{}
	for {
		token, err := dec.Token()
		if err != nil {
			// Quit for-loop when EOF is reached
			if err == io.EOF {
				break
			}
			return err
		}

		sToken, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		token, _ = dec.Token()
		char, ok := token.(xml.CharData)
		if !ok {
			continue
		}

		count, err := strconv.Atoi(string(char))
		if err != nil {
			return err
		}

		(*c)[sToken.Name.Local] = count

	}

	return nil
}

func (c IntList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: start.Name})
	for k, v := range c {
		e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: k}})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return e.Flush()
}

type Forms []Form

type Form struct {
	Type              string  `json:"Type"`
	DataFormatId      string  `json:"DataFormatId"`
	DataFormatVersion string  `json:"DataFormatVersion"`
	FormData          Melding `json:"FormData"`
}
