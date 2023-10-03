package geo

import (
	"embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed assets/maximind/*.mmdb
var embedFS embed.FS

const (
	maxmindCountryTestDB string = "assets/maximind/GeoIP2-Country-Test.mmdb" // test database, only covers NA, EU, AS continents
)

type countryLookupTest struct {
	ip string

	expectedCountryCode       string
	expectedCountryAlpha3Code string

	expectedContinentCode string

	expectedEmpty bool
	expectedError bool
}

func TestCountryLookupWithMaxmind(t *testing.T) {

	file, err := embedFS.Open(maxmindCountryTestDB)
	if err != nil {
		t.Fatal(fmt.Errorf("embedFS.Open: %w", err))
	}

	maximind, err := NewMaxMindReader(file)
	if err != nil {
		t.Fatal(fmt.Errorf("NewMaxMindn: %w", err))
	}

	ipTests := []countryLookupTest{
		{
			ip:                        "::149.101.100.0",
			expectedCountryCode:       "US",
			expectedCountryAlpha3Code: "USA",
			expectedContinentCode:     "NA",
			expectedEmpty:             false,
			expectedError:             false,
		},
		{
			ip:                        "81.2.69.142",
			expectedCountryCode:       "GB",
			expectedCountryAlpha3Code: "GBR",
			expectedContinentCode:     "EU",
			expectedEmpty:             false,
			expectedError:             false,
		},

		{
			ip:                        "2001:218::",
			expectedCountryCode:       "JP",
			expectedCountryAlpha3Code: "JPN",
			expectedContinentCode:     "AS",
			expectedEmpty:             false,
			expectedError:             false,
		},

		{
			ip:                        "1", // wrong ip
			expectedCountryCode:       "",
			expectedContinentCode:     "",
			expectedCountryAlpha3Code: "",
			expectedEmpty:             false,
			expectedError:             true,
		},

		{
			ip:                        "127.0.0.1", // local ip
			expectedCountryCode:       "ZZ",
			expectedContinentCode:     "ZZ",
			expectedCountryAlpha3Code: "ZZZ",
			expectedEmpty:             true,
			expectedError:             false,
		},
	}

	for _, test := range ipTests {
		country, err := maximind.CountryByIPString(test.ip)

		if test.expectedError {
			assert.NotNil(t, err, "expected error")
			assert.Nil(t, country, "expected country to be nil ")
			continue
		}

		assert.Nil(t, err, "expected no error")
		assert.NotNil(t, country, "expected country to not be nil ")

		if test.expectedEmpty {
			assert.Equal(t, true, country.IsEmpty(), "expected country is empty")
		} else {
			assert.Equal(t, false, country.IsEmpty(), "expected country is not empty")
		}

		assert.Equal(t, test.expectedCountryCode, country.CountryAlpha2Code(), "mismatch country code")
		assert.Equal(t, test.expectedCountryAlpha3Code, country.CountryAlpha3Code(), "mismatch country alpha3 code")
		assert.Equal(t, test.expectedContinentCode, country.ContinentCode(), "mismatch contininent code")
	}
}