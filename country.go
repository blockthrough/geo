package geo

import (
	"github.com/oschwald/geoip2-golang"
)

// Country - a type definition on geoip2.Country data struct while providing heloer functions to retrieve certain data in convinent way and additional country data maxmind db does not provide
// user is expected not to directly modify it.
type Country geoip2.Country

// CountryAlpha2Code - a helper function to retrieve 2-letter ISO code for country from maxmind DB, if the country is found, the code will be "ZZ"
func (c Country) CountryAlpha2Code() string {
	if c.isCountryAlpha2CodeValid() {
		return c.Country.IsoCode
	}

	return UnknownAlpha2Code
}

// ContinentCode - a helper function to retrieve 2-letter ISO code for country from maxmind DB
func (c Country) ContinentCode() string {
	if c.isCountryAlpha2CodeValid() {
		return c.Continent.Code
	}

	return UnknownAlpha2Code
}

// CountryAlpha3Code - return 3-letter ISO code for country
func (c Country) CountryAlpha3Code() string {
	if c.isCountryAlpha2CodeValid() {
		return CountryAlpha2CodeToAlpha3Code(c.CountryAlpha2Code())
	}

	return UnknownAlpha3Code
}

// IsUnknown - helper function to determine if the country is unknown
func (c Country) IsUnknown() bool {
	return c.CountryAlpha2Code() == UnknownAlpha2Code
}

// isEmpty - check if the country has an id associated with MaxMindDB, true means a valid entry, false means MaxMindDB does not find it
func (c Country) isEmpty() bool {
	return c.Country.GeoNameID == 0
}

func (c Country) isCountryAlpha2CodeValid() bool {
	return !c.isEmpty() && len(c.Country.IsoCode) == 2
}
