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
						GJD_ALBANIA             string `xml:"GJD_ALBANIA"`
						GJD_ANDORRA             string `xml:"GJD_ANDORRA"`
						GJD_BELGIA              string `xml:"GJD_BELGIA"`
						GJD_BOSNIAOGHERZEGOVINA string `xml:"GJD_BOSNIAOGHERZEGOVINA"`
						GJD_BULGARIA            string `xml:"GJD_BULGARIA"`
						GJD_ESTLAND             string `xml:"GJD_ESTLAND"`
						GJD_FRANKRIKE           string `xml:"GJD_FRANKRIKE"`
						GJD_HELLAS              string `xml:"GJD_HELLAS"`
						GJD_HVITERUSSLAND       string `xml:"GJD_HVITERUSSLAND"`
						GJD_IRLAND              string `xml:"GJD_IRLAND"`
						GJD_ITALIA              string `xml:"GJD_ITALIA"`
						GJD_KOSOVO              string `xml:"GJD_KOSOVO"`
						GJD_KROATIA             string `xml:"GJD_KROATIA"`
						GJD_KYPROS              string `xml:"GJD_KYPROS"`
						GJD_LATVIA              string `xml:"GJD_LATVIA"`
						GJD_LIECHTENSTEIN       string `xml:"GJD_LIECHTENSTEIN"`
						GJD_LITAUEN             string `xml:"GJD_LITAUEN"`
						GJD_LUXEMBOURG          string `xml:"GJD_LUXEMBOURG"`
						GJD_MAKEDONIA           string `xml:"GJD_MAKEDONIA"`
						GJD_MALTA               string `xml:"GJD_MALTA"`
						GJD_MOLDOVA             string `xml:"GJD_MOLDOVA"`
						GJD_MONACO              string `xml:"GJD_MONACO"`
						GJD_MONTNEGRO           string `xml:"GJD_MONTNEGRO"`
						GJD_NEDERLAND           string `xml:"GJD_NEDERLAND"`
						GJD_POLEN               string `xml:"GJD_POLEN"`
						GJD_PORTUGAL            string `xml:"GJD_PORTUGAL"`
						GJD_ROMANIA             string `xml:"GJD_ROMANIA"`
						GJD_RUSSLAND            string `xml:"GJD_RUSSLAND"`
						GJD_SANMARINO           string `xml:"GJD_SANMARINO"`
						GJD_SERBIA              string `xml:"GJD_SERBIA"`
						GJD_SLOVAKIA            string `xml:"GJD_SLOVAKIA"`
						GJD_SLOVENIA            string `xml:"GJD_SLOVENIA"`
						GJD_SPANIA              string `xml:"GJD_SPANIA"`
						GJD_STORBRITANNIA       string `xml:"GJD_STORBRITANNIA"`
						GJD_SVEITS              string `xml:"GJD_SVEITS"`
						GJD_TSJEKKIA            string `xml:"GJD_TSJEKKIA"`
						GJD_TYRKIA              string `xml:"GJD_TYRKIA"`
						GJD_TYSKLAND            string `xml:"GJD_TYSKLAND"`
						GJD_UKRAINA             string `xml:"GJD_UKRAINA"`
						GJD_UNGARN              string `xml:"GJD_UNGARN"`
						GJD_VATIKANSTATEN       string `xml:"GJD_VATIKANSTATEN"`
						GJD_OSTERRIKE           string `xml:"GJD_OSTERRIKE"`
					} `xml:"europa"`
					Asia struct {
						GJD_EMIRATER    string `xml:"GJD_EMIRATER"`
						GJD_INDIA       string `xml:"GJD_INDIA"`
						GJD_INDONESIA   string `xml:"GJD_INDONESIA"`
						GJD_ISRAEL      string `xml:"GJD_ISRAEL"`
						GJD_JAPAN       string `xml:"GJD_JAPAN"`
						GJD_KINA        string `xml:"GJD_KINA"`
						GJD_MALAYSIA    string `xml:"GJD_MALAYSIA"`
						GJD_QATAR       string `xml:"GJD_QATAR"`
						GJD_SINGAPORE   string `xml:"GJD_SINGAPORE"`
						GJD_SOR_KOREA   string `xml:"GJD_SOR_KOREA"`
						GJD_TAIWAN      string `xml:"GJD_TAIWAN"`
						GJD_THAILAND    string `xml:"GJD_THAILAND"`
						GJD_ASIA_ELLERS string `xml:"GJD_ASIA_ELLERS"`
					} `xml:"asia"`
					Amerika struct {
						GJD_BRASIL               string `xml:"GJD_BRASIL"`
						GJD_CANADA               string `xml:"GJD_CANADA"`
						GJD_MEXICO               string `xml:"GJD_MEXICO"`
						GJD_USA                  string `xml:"GJD_USA"`
						GJD_LATIN_AMERIKA_ELLERS string `xml:"GJD_LATIN_AMERIKA_ELLERS"`
					} `xml:"amerika"`
					Oceania struct {
						GJD_AUSTRALIA      string `xml:"GJD_AUSTRALIA"`
						GJD_OCEANIA_ELLERS string `xml:"GJD_OCEANIA_ELLERS"`
					} `xml:"oceania"`
					Afrika struct {
						GJD_SOR_AFRIKA    string `xml:"GJD_SOR_AFRIKA"`
						GJD_AFRIKA_ELLERS string `xml:"GJD_AFRIKA_ELLERS"`
					} `xml:"afrika"`
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
