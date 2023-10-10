package geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlpha2ToAlpha3CountryCode(t *testing.T) {

	type testCase struct {
		alpha2CountryCode         string
		expectedAlpha3CountryCode string
	}

	tests := []testCase{
		{
			alpha2CountryCode:         "A",
			expectedAlpha3CountryCode: "ZZZ",
		},
		{
			alpha2CountryCode:         "US",
			expectedAlpha3CountryCode: "USA",
		},

		{
			alpha2CountryCode:         "ZZ",
			expectedAlpha3CountryCode: "ZZZ",
		},

		{
			alpha2CountryCode:         "GB",
			expectedAlpha3CountryCode: "GBR",
		},
		{
			alpha2CountryCode:         "CA",
			expectedAlpha3CountryCode: "CAN",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAlpha3CountryCode, CountryAlpha2CodeToAlpha3Code(test.alpha2CountryCode), "mismatch alpha3 country code")
	}
}

func TestAlpha3ToAlpha2CountryCode(t *testing.T) {

	type testCase struct {
		alpha3CountryCode         string
		expectedAlpha2CountryCode string
	}

	tests := []testCase{
		{
			alpha3CountryCode:         "ZZZ",
			expectedAlpha2CountryCode: "ZZ",
		},
		{
			alpha3CountryCode:         "A",
			expectedAlpha2CountryCode: "ZZ",
		},
		{
			alpha3CountryCode:         "USA",
			expectedAlpha2CountryCode: "US",
		},

		{
			alpha3CountryCode:         "CAN",
			expectedAlpha2CountryCode: "CA",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedAlpha2CountryCode, CountryAlpha3CodeToAlpha2Code(test.expectedAlpha2CountryCode), "mismatch alpha3 country code")
	}
}
