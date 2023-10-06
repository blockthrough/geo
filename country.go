package geo

import (
	"github.com/oschwald/geoip2-golang"
)

// Country - a type definition on geoip2.Country data struct while providing helper functions to retrieve certain data in convinent way and additional country data maxmind db does not provide
// user is expected not to directly modify it.
type country geoip2.Country

// CountryAlpha2Code - a helper function to retrieve 2-letter ISO code for country from maxmind DB, if the country is unknown, the code will be "ZZ"
func (c country) CountryAlpha2Code() string {
	if c.isCountryAlpha2CodeValid() {
		return c.Country.IsoCode
	}

	return UnknownAlpha2Code
}

// ContinentCode - a helper function to retrieve 2-letter ISO code for country from maxmind DB, if the country is known, the code will be "ZZ"
func (c country) ContinentCode() string {
	if c.isCountryAlpha2CodeValid() {
		return c.Continent.Code
	}

	return UnknownAlpha2Code
}

// CountryAlpha3Code - return 3-letter ISO code for country, if the country is unknown, the code will be "ZZZ"
func (c country) CountryAlpha3Code() string {
	if c.isCountryAlpha2CodeValid() {
		return CountryAlpha2CodeToAlpha3Code(c.CountryAlpha2Code())
	}

	return UnknownAlpha3Code
}

// IsUnknown - helper function to determine if the country is unknown
func (c country) IsUnknown() bool {
	return c.CountryAlpha2Code() == UnknownAlpha2Code
}

// isEmpty - check if the country has an id associated with MaxMindDB, true means a valid entry, false means no valid entry
func (c country) isEmpty() bool {
	return c.Country.GeoNameID == 0
}

func (c country) isCountryAlpha2CodeValid() bool {
	return !c.isEmpty() && len(c.Country.IsoCode) == 2
}
