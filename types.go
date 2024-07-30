package altinn

import (
	"encoding/xml"
	"io"
	"strconv"
)

type Melding struct {
	XMLName            xml.Name `xml:"melding"`
	Seres              string   `xml:"xmlns:seres,attr"`
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
					Norden struct {
						GJD_DANMARK int `xml:"GJD_DANMARK"`
						GJD_FINLAND int `xml:"GJD_FINLAND"`
						GJD_ISLAND  int `xml:"GJD_ISLAND"`
						GJD_NORGE   int `xml:"GJD_NORGE"`
						GJD_SVERIGE int `xml:"GJD_SVERIGE"`
					} `xml:"norden"`
					Europa struct {
						GJD_ALBANIA             int `xml:"GJD_ALBANIA,omitempty"`
						GJD_ANDORRA             int `xml:"GJD_ANDORRA,omitempty"`
						GJD_BELGIA              int `xml:"GJD_BELGIA,omitempty"`
						GJD_BOSNIAOGHERZEGOVINA int `xml:"GJD_BOSNIAOGHERZEGOVINA,omitempty"`
						GJD_BULGARIA            int `xml:"GJD_BULGARIA,omitempty"`
						GJD_ESTLAND             int `xml:"GJD_ESTLAND,omitempty"`
						GJD_FRANKRIKE           int `xml:"GJD_FRANKRIKE,omitempty"`
						GJD_HELLAS              int `xml:"GJD_HELLAS,omitempty"`
						GJD_HVITERUSSLAND       int `xml:"GJD_HVITERUSSLAND,omitempty"`
						GJD_IRLAND              int `xml:"GJD_IRLAND,omitempty"`
						GJD_ITALIA              int `xml:"GJD_ITALIA,omitempty"`
						GJD_KOSOVO              int `xml:"GJD_KOSOVO,omitempty"`
						GJD_KROATIA             int `xml:"GJD_KROATIA,omitempty"`
						GJD_KYPROS              int `xml:"GJD_KYPROS,omitempty"`
						GJD_LATVIA              int `xml:"GJD_LATVIA,omitempty"`
						GJD_LIECHTENSTEIN       int `xml:"GJD_LIECHTENSTEIN,omitempty"`
						GJD_LITAUEN             int `xml:"GJD_LITAUEN,omitempty"`
						GJD_LUXEMBOURG          int `xml:"GJD_LUXEMBOURG,omitempty"`
						GJD_MAKEDONIA           int `xml:"GJD_MAKEDONIA,omitempty"`
						GJD_MALTA               int `xml:"GJD_MALTA,omitempty"`
						GJD_MOLDOVA             int `xml:"GJD_MOLDOVA,omitempty"`
						GJD_MONACO              int `xml:"GJD_MONACO,omitempty"`
						GJD_MONTNEGRO           int `xml:"GJD_MONTNEGRO,omitempty"`
						GJD_NEDERLAND           int `xml:"GJD_NEDERLAND,omitempty"`
						GJD_POLEN               int `xml:"GJD_POLEN,omitempty"`
						GJD_PORTUGAL            int `xml:"GJD_PORTUGAL,omitempty"`
						GJD_ROMANIA             int `xml:"GJD_ROMANIA,omitempty"`
						GJD_RUSSLAND            int `xml:"GJD_RUSSLAND,omitempty"`
						GJD_SANMARINO           int `xml:"GJD_SANMARINO,omitempty"`
						GJD_SERBIA              int `xml:"GJD_SERBIA,omitempty"`
						GJD_SLOVAKIA            int `xml:"GJD_SLOVAKIA,omitempty"`
						GJD_SLOVENIA            int `xml:"GJD_SLOVENIA,omitempty"`
						GJD_SPANIA              int `xml:"GJD_SPANIA,omitempty"`
						GJD_STORBRITANNIA       int `xml:"GJD_STORBRITANNIA,omitempty"`
						GJD_SVEITS              int `xml:"GJD_SVEITS,omitempty"`
						GJD_TSJEKKIA            int `xml:"GJD_TSJEKKIA,omitempty"`
						GJD_TYRKIA              int `xml:"GJD_TYRKIA,omitempty"`
						GJD_TYSKLAND            int `xml:"GJD_TYSKLAND,omitempty"`
						GJD_UKRAINA             int `xml:"GJD_UKRAINA,omitempty"`
						GJD_UNGARN              int `xml:"GJD_UNGARN,omitempty"`
						GJD_VATIKANSTATEN       int `xml:"GJD_VATIKANSTATEN,omitempty"`
						GJD_OSTERRIKE           int `xml:"GJD_OSTERRIKE,omitempty"`
					} `xml:"europa,omitempty"`
					Asia struct {
						GJD_EMIRATER    int `xml:"GJD_EMIRATER,omitempty"`
						GJD_INDIA       int `xml:"GJD_INDIA,omitempty"`
						GJD_INDONESIA   int `xml:"GJD_INDONESIA,omitempty"`
						GJD_ISRAEL      int `xml:"GJD_ISRAEL,omitempty"`
						GJD_JAPAN       int `xml:"GJD_JAPAN,omitempty"`
						GJD_KINA        int `xml:"GJD_KINA,omitempty"`
						GJD_MALAYSIA    int `xml:"GJD_MALAYSIA,omitempty"`
						GJD_QATAR       int `xml:"GJD_QATAR,omitempty"`
						GJD_SINGAPORE   int `xml:"GJD_SINGAPORE,omitempty"`
						GJD_SOR_KOREA   int `xml:"GJD_SOR_KOREA,omitempty"`
						GJD_TAIWAN      int `xml:"GJD_TAIWAN,omitempty"`
						GJD_THAILAND    int `xml:"GJD_THAILAND,omitempty"`
						GJD_ASIA_ELLERS int `xml:"GJD_ASIA_ELLERS,omitempty"`
					} `xml:"asia,omitempty"`
					Amerika struct {
						GJD_BRASIL               int `xml:"GJD_BRASIL,omitempty"`
						GJD_CANADA               int `xml:"GJD_CANADA,omitempty"`
						GJD_MEXICO               int `xml:"GJD_MEXICO,omitempty"`
						GJD_USA                  int `xml:"GJD_USA,omitempty"`
						GJD_LATIN_AMERIKA_ELLERS int `xml:"GJD_LATIN_AMERIKA_ELLERS,omitempty"`
					} `xml:"amerika,omitempty"`
					Oceania struct {
						GJD_AUSTRALIA      int `xml:"GJD_AUSTRALIA,omitempty"`
						GJD_OCEANIA_ELLERS int `xml:"GJD_OCEANIA_ELLERS,omitempty"`
					} `xml:"oceania,omitempty"`
					Afrika struct {
						GJD_SOR_AFRIKA    int `xml:"GJD_SOR_AFRIKA,omitempty"`
						GJD_AFRIKA_ELLERS int `xml:"GJD_AFRIKA_ELLERS,omitempty"`
					} `xml:"afrika,omitempty"`
				} `xml:"land"`
			} `xml:"beleggHotell"`
		} `xml:"belegg"`
	} `xml:"Skjemadata"`
}

func (m Melding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	m.Seres = "http://seres.no/xsd/forvaltningsdata"
	type Alias Melding
	return e.Encode(Alias(m))
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
