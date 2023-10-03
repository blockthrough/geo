package geo

import (
	"testing"

	"github.com/oschwald/geoip2-golang"
	"github.com/stretchr/testify/assert"
)

type countryTestCase struct {
	country Country

	expectedCountryCode       string
	expectedCountryAlpha3Code string
	expectedContinentCode     string
	expectedUnknown           bool
}

func TestCountry(t *testing.T) {

	tests := []countryTestCase{
		{
			country:                   Country(geoip2.Country{}),
			expectedCountryCode:       "ZZ",
			expectedCountryAlpha3Code: "ZZZ",
			expectedContinentCode:     "ZZ",

			expectedUnknown: true,
		},
		{
			country: Country(geoip2.Country{
				Country: struct {
					Names             map[string]string "maxminddb:\"names\""
					IsoCode           string            "maxminddb:\"iso_code\""
					GeoNameID         uint              "maxminddb:\"geoname_id\""
					IsInEuropeanUnion bool              "maxminddb:\"is_in_european_union\""
				}{
					GeoNameID: 1,
					IsoCode:   "US",
				},
				Continent: struct {
					Names     map[string]string "maxminddb:\"names\""
					Code      string            "maxminddb:\"code\""
					GeoNameID uint              "maxminddb:\"geoname_id\""
				}{Code: "NA"},
			}),
			expectedCountryCode:       "US",
			expectedCountryAlpha3Code: "USA",
			expectedContinentCode:     "NA",
			expectedUnknown:           false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedCountryCode, test.country.CountryAlpha2Code(), "mismatch country code")
		assert.Equal(t, test.expectedCountryAlpha3Code, test.country.CountryAlpha3Code(), "mismatch country alpha3 code")
		assert.Equal(t, test.expectedContinentCode, test.country.ContinentCode(), "mismatch continent code")
		assert.Equal(t, test.expectedUnknown, test.country.IsUnknown(), "mismatch unknown expectation")
	}

}
