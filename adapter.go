package geo

import "github.com/biter777/countries"

const UnknownAlpha2Code = "ZZ"  // ZZ is the commonly recoginized country/continient code for unknown, as specified in ISO alpha2
const UnknownAlpha3Code = "ZZZ" // ZZZ is derived from "ZZ" to represent 3 letter country/continent code for unknown, as specified in ISO alpha3

// CountryAlpha2CodeToAlpha3Code - get a 3-letter country code if a country identified by 2-letter country code
func CountryAlpha2CodeToAlpha3Code(alpha2Code string) string {
	alpha3 := countries.ByName(alpha2Code).Alpha3()
	if alpha3 == "Unknown" {
		return UnknownAlpha3Code
	}

	return alpha3
}

// CountryAlpha3CodeToAlpha2Code - get a 2-letter country code if a country identified by 3-letter country code
func CountryAlpha3CodeToAlpha2Code(alpha3Code string) string {
	alpha2 := countries.ByName(alpha3Code).Alpha2()
	if alpha2 == "Unknown" {
		return UnknownAlpha2Code
	}

	return alpha2
}
