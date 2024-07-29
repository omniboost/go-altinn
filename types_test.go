package altinn_test

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"testing"

	altinn "github.com/omniboost/go-altinn"
)

func TestMarshalMelding(t *testing.T) {
	b := []byte(`
<?xml version='1.0' encoding='utf-8'?>
<melding xmlns:seres= "http://seres.no/xsd/forvaltningsdata" dataFormatProvider="SERES" dataFormatId="6400" dataFormatVersion="45188">
  <InternInformasjon>
    <periodeFritekst>juni 2024</periodeFritekst>
    <rapportPeriode>
      <fomDato>2024-06-01</fomDato>
      <tomDato>2024-06-30</tomDato>
    </rapportPeriode>
    <raNummer>RA-0297</raNummer>
    <delRegNummer>0</delRegNummer>
    <identnummerEnhet>973068075</identnummerEnhet>
    <sendtFraSluttbrukersystem>true</sendtFraSluttbrukersystem>
    <statistiskEnhet>
      <enhetsident>973068075</enhetsident>
      <enhetstype>BEDR</enhetstype>
    </statistiskEnhet>
    <skjemaidentifikasjon>
      <periodenummer>06</periodenummer>
      <periodetype>MND</periodetype>
      <periodeAAr>2024</periodeAAr>
      <undersoekelsesnummer>78</undersoekelsesnummer>
    </skjemaidentifikasjon>
  </InternInformasjon>
  <KontaktpersonOgKommentarfelt>
    <kontaktperson>
      <epostadresse>kristin.frigstad@strawberry.no</epostadresse>
      <navn>Kristin Frigstad</navn>
      <telefonSFU>38128600</telefonSFU>
    </kontaktperson>
  </KontaktpersonOgKommentarfelt>
  <ForetakOgVirksomhetsopplysninger>
    <virksomhet>
      <organisasjonsnummerVirksomhet>973068075</organisasjonsnummerVirksomhet>
      <navnVirksomhet>Clarion Hotel Ernst</navnVirksomhet>
      <adresseVirksomhet>
        <gateadresse>Radhusgaten 2</gateadresse>
        <postnummer>4664</postnummer>
        <poststed>Kristiansand</poststed>
      </adresseVirksomhet>
    </virksomhet>
  </ForetakOgVirksomhetsopplysninger>
  <StatusVirksomhetMedDriftsperiode>
    <driftsstatusPeriode>jaDrift</driftsstatusPeriode>
  </StatusVirksomhetMedDriftsperiode>
  <Naeringskontrollspoersmaal>
    <visNaeringskontrollJaNeiPrefill />
    <Naeringskontroll>
      <naeringskode>55.101</naeringskode>
      <naeringstekst>Drift av hoteller, pensjonater og moteller med restaurant</naeringstekst>
      <naeringsbeskrivelse />
      <nyNaeringsbeskrivelse>
        <alltidViktigsteAktivitet />
      </nyNaeringsbeskrivelse>
    </Naeringskontroll>
  </Naeringskontrollspoersmaal>
  <Oppgavebyrde>
    <visOppgavebyrdeJaNeiPrefill>2</visOppgavebyrdeJaNeiPrefill>
    <tidsbrukSamleInformasjon />
    <tidsbrukSelveSkjemautfyllingen />
    <tidsbrukTotalHjelpAvAndre />
  </Oppgavebyrde>
  <Skjemadata>
    <rapporteringsenhet>
      <typeHotell />
      <typeVandrehjem />
      <typeHytter />
      <typeCamping />
      <typeAnnen />
    </rapporteringsenhet>
    <kapasitet>
      <kapasitetHotell>
        <antDager>
          <endret>1</endret>
          <antallEndret>29</antallEndret>
        </antDager>
        <antGjesterom>
          <endret>1</endret>
          <antallEndret>199</antallEndret>
        </antGjesterom>
        <antRullestol>
          <endret>1</endret>
          <antallEndret>11</antallEndret>
        </antRullestol>
        <antFasteSengeplasser>
          <endret>1</endret>
          <antallEndret>378</antallEndret>
        </antFasteSengeplasser>
      </kapasitetHotell>
      <kapasitetHytte>
        <antDager />
        <antUtleieenheter />
        <antFasteSengeplasser />
      </kapasitetHytte>
      <kapasitetCamping>
        <antDager />
        <antUtekapasitet />
        <antInnekapasitet />
        <antFasteSengeplasser />
      </kapasitetCamping>
    </kapasitet>
    <omsetningHotell>
      <romdoegn>4833</romdoegn>
      <losjiomsetning>7907433</losjiomsetning>
    </omsetningHotell>
    <belegg>
      <beleggHotell>
        <antGjester>4268</antGjester>
        <antNorskeGjester>3623</antNorskeGjester>
        <antGjestedoegn>6673</antGjestedoegn>
        <ferieFritidGjestedoegn>1568</ferieFritidGjestedoegn>
        <kursGjestedoegn>338</kursGjestedoegn>
        <forretningsGjestedoegn>4767</forretningsGjestedoegn>
        <land>
          <norden>
            <GJD_DANMARK>48</GJD_DANMARK>
            <GJD_FINLAND>13</GJD_FINLAND>
            <GJD_ISLAND>4</GJD_ISLAND>
            <GJD_NORGE>5741</GJD_NORGE>
            <GJD_SVERIGE>102</GJD_SVERIGE>
          </norden>
          <europa>
            <GJD_BELGIA>8</GJD_BELGIA>
            <GJD_FRANKRIKE>6</GJD_FRANKRIKE>
            <GJD_ITALIA>9</GJD_ITALIA>
            <GJD_KROATIA>3</GJD_KROATIA>
            <GJD_NEDERLAND>19</GJD_NEDERLAND>
            <GJD_POLEN>2</GJD_POLEN>
            <GJD_SLOVAKIA>3</GJD_SLOVAKIA>
            <GJD_STORBRITANNIA>22</GJD_STORBRITANNIA>
            <GJD_SVEITS>20</GJD_SVEITS>
            <GJD_TYSKLAND>70</GJD_TYSKLAND>
          </europa>
          <asia>
            <GJD_JAPAN>2</GJD_JAPAN>
            <GJD_ASIA_ELLERS>7</GJD_ASIA_ELLERS>
          </asia>
          <amerika>
            <GJD_BRASIL>4</GJD_BRASIL>
            <GJD_USA>584</GJD_USA>
          </amerika>
          <oceania>
            <GJD_AUSTRALIA>6</GJD_AUSTRALIA>
          </oceania>
          <afrika />
        </land>
      </beleggHotell>
    </belegg>
  </Skjemadata>
</melding>`)

	melding := altinn.Melding{}

	// convert xml to struct
	err := xml.Unmarshal(b, &melding)
	if err != nil {
		t.Error(err)
	}

	// print out struct as xml
	b, err = xml.MarshalIndent(melding, "", "  ")
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))

	// convert struct to json
	b, err = json.MarshalIndent(melding, "", "  ")
	if err != nil {
		t.Error(err)
	}

	// convert json back to struct
	err = json.Unmarshal(b, &melding)
	if err != nil {
		t.Error(err)
	}

	// print out struct as xml
	b, err = xml.MarshalIndent(melding, "", "  ")
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
}
