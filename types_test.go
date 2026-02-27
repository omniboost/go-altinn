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
<A3_RS0297_M
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:xsd="http://www.w3.org/2001/XMLSchema" dataFormatProvider="ALTINNSTUDIO" dataFormatId="A31000" dataFormatVersion="1">
    <InternInfo>
        <periodeNummer>01</periodeNummer>
        <periodeAAr>2026</periodeAAr>
        <enhetsOrgNr>973222597</enhetsOrgNr>
    </InternInfo>
    <SkjemaData>
        <sluttbrukerSystemNavn>Mews</sluttbrukerSystemNavn>
        <iDriftJaNei>1</iDriftJaNei>
        <antallDagerAApen>31</antallDagerAApen>
        <tilgjengeligeGjesterom>144</tilgjengeligeGjesterom>
        <tilgjengeligeRomRullestol>0</tilgjengeligeRomRullestol>
        <fasteSengeplasser>144</fasteSengeplasser>
        <romdogn>1400</romdogn>
        <losjiomsetning>1714831</losjiomsetning>
        <antallNorskeGjester>1212</antallNorskeGjester>
        <antallUtlGjester>104</antallUtlGjester>
        <antallGjestedogn>1864</antallGjestedogn>
        <bostedsLand>
            <bostedsLandID>000</bostedsLandID>
            <antallGjestedognLand>1648</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>101</bostedsLandID>
            <antallGjestedognLand>24</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>106</bostedsLandID>
            <antallGjestedognLand>19</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>122</bostedsLandID>
            <antallGjestedognLand>1</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>121</bostedsLandID>
            <antallGjestedognLand>4</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>124</bostedsLandID>
            <antallGjestedognLand>4</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>136</bostedsLandID>
            <antallGjestedognLand>3</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>127</bostedsLandID>
            <antallGjestedognLand>18</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>131</bostedsLandID>
            <antallGjestedognLand>5</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>132</bostedsLandID>
            <antallGjestedognLand>6</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>133</bostedsLandID>
            <antallGjestedognLand>3</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>137</bostedsLandID>
            <antallGjestedognLand>36</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>139</bostedsLandID>
            <antallGjestedognLand>29</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>141</bostedsLandID>
            <antallGjestedognLand>20</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>143</bostedsLandID>
            <antallGjestedognLand>1</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>144</bostedsLandID>
            <antallGjestedognLand>7</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>444</bostedsLandID>
            <antallGjestedognLand>3</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>448</bostedsLandID>
            <antallGjestedognLand>6</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>612</bostedsLandID>
            <antallGjestedognLand>15</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>684</bostedsLandID>
            <antallGjestedognLand>8</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>805</bostedsLandID>
            <antallGjestedognLand>4</antallGjestedognLand>
        </bostedsLand>
        <bostedsLand>
            <bostedsLandID>990</bostedsLandID>
            <antallGjestedognLand>0</antallGjestedognLand>
        </bostedsLand>
        <oppholdsFormaal>
            <formaalID>1</formaalID>
            <antallGjestedognFormaal>1223</antallGjestedognFormaal>
        </oppholdsFormaal>
        <oppholdsFormaal>
            <formaalID>2</formaalID>
            <antallGjestedognFormaal>0</antallGjestedognFormaal>
        </oppholdsFormaal>
        <oppholdsFormaal>
            <formaalID>3</formaalID>
            <antallGjestedognFormaal>641</antallGjestedognFormaal>
        </oppholdsFormaal>
    </SkjemaData>
</A3_RS0297_M>
`)

	melding := altinn.RS0297FormData{}

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
