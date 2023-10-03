package geo

import (
	"github.com/oschwald/geoip2-golang"
)

// Country - a type definition on geoip2.Country data struct while providing heloer functions to retrieve certain data in convinent way and additional country data maxmind db does not provide
// user is expected not to directly modify it
type Country geoip2.Country

// CountryAlpha2Code - a helper function to retrieve 2-letter ISO code for country from maxmind DB, if the country is found, the code will be "ZZ"
func (c Country) CountryAlpha2Code() string {
	if c.isEmpty() || len(c.Country.IsoCode) != 2 {
		return UnknownAlpha2Code
	}

	return c.Country.IsoCode
}

// ContinentCode - a helper function to retrieve 2-letter ISO code for country from maxmind DB
func (c Country) ContinentCode() string {
	if c.isEmpty() || len(c.Country.IsoCode) != 2 {
		return UnknownAlpha2Code
	}

	return c.Continent.Code
}

// CountryAlpha3Code - return 3-letter ISO code for country
func (c Country) CountryAlpha3Code() string {
	if c.isEmpty() || len(c.Country.IsoCode) != 2 {
		return UnknownAlpha3Code
	}

	return CountryAlpha2CodeToAlpha3Code(c.CountryAlpha2Code())
}

// IsUnknown - helper function to determine if the country is unknown
func (c Country) IsUnknown() bool {
	return c.CountryAlpha2Code() == UnknownAlpha2Code
}

// isEmpty - check if the country has an id associated with MaxMindDB, true means a valid entry, false means MaxMindDB does not find it
func (c Country) isEmpty() bool {
	return c.Country.GeoNameID == 0
}
