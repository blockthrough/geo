package geo

import (
	"github.com/biter777/countries"
	"github.com/oschwald/geoip2-golang"
)

const UnknownAlpha2Code = "ZZ"  // ZZ is the commonly recoginized country/continient code for unknown, as specified in ISO alpha2
const UnknownAlpha3Code = "ZZZ" // ZZZ is derived from "ZZ" to represent 3 letter country/continent code for unknown, as specified in ISO alpha3

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

	alpha3 := countries.ByName(c.CountryAlpha2Code()).Alpha3()
	if alpha3 == "Unknown" {
		return UnknownAlpha3Code
	}

	return alpha3
}

// IsUnknown - helper function to determine if the country is unknown
func (c Country) IsUnknown() bool {
	return c.CountryAlpha2Code() == UnknownAlpha2Code
}

// isEmpty - check if the country has an id associated with MaxMindDB, true means a valid entry, false means MaxMindDB does not find it
func (c Country) isEmpty() bool {
	return c.Country.GeoNameID == 0
}
