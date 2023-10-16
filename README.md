# Geo Information Library For Go
[![CircleCI](https://circleci.com/gh/blockthrough/geo.svg?style=svg&circle-token=b0554d26f90621f9996755fe9fd6665e74cabcbe)](<https://app.circleci.com/pipelines/github/blockthrough/geo?branch=master>) [![codecov](https://codecov.io/gh/blockthrough/geo/branch/master/graph/badge.svg?token=DK6KZBFHML)](https://codecov.io/gh/blockthrough/geo)


The library currently supports reads MaxMind GeoIP2 database and provide a thin decorator/helper function on the returned data for easy access and decoration on top of the geo info provided by MaxMind. It also provides a normalized way to represent the geo info in standard ISO format the scenario of unknown.

The implementation of this library is based upon other 2 open source repositories:

1. https://github.com/biter777/countries 
2. https://github.com/oschwald/geoip2-golang


## Goal
The library is visioned to support more features related to processing of the geo info data in the future. That means, support reading from multiple geo database providers and expose more helpers on the data for commonly encountered user scenarios.


## Features
1. Support easy lookup for geo information from MaxMind DB using IP string.
2. Support both alpha2 country code and alpha3 country code.
3. Support any `io.Reader` for user to provide the MaxMind `.mmdb` file.
4. Support easy access of the database metadata info like build time and version.
4. Provide default return value of `ZZ` and `ZZZ` as specified in ISO standard if the country could not be found by the input IP using helper functions.


## Installation
```
go get github.com/blockthrough/geo@v1.0.0
```

## How To Use

### Read Maxmind DB and LookUp Geo With IP
```go
package main 

import (
    "embed"
    "fmt"
    "github.com/blockthrough/geo"
)

// go:embed <your_maxmind_db_file_path>
var embedFS embed.FS

func main() {
	file, err := embedFS.Open("<your_maxmind_db_file_path>")
	if err != nil {
		return
	}

	maxmind, err := geo.NewMaxMindFromReader(file)
	if err != nil {
		return
	}

	// check the meta information of the database
	fmt.Sprintf("test db: %t", maxmind.IsTestDB())
	fmt.Sprintf("db build time: %s", maxmind.BuildTimestamp())
	fmt.Sprintf("db version: %s", maxmind.Version())

	// you can pass ipv4 or ipv6  address
	country, err := maxmind.CountryByIPString("127.0.0.1")
	if err != nil {
		return 
	}

	fmt.Sprintf("unknown country: %s", country.IsUnknown())               // is the country unknown?
	fmt.Sprintf("country code: %s", country.CountryAlpha2Code())          // 2-letter country code
	fmt.Sprintf("country 3-letter code: %s", country.CountryAlpha3Code()) // 3-letter country code
	fmt.Sprintf("continent code: %s", country.ContinentCode())            // 2-letter continent code

	// you can also use adapter function directly if you want to get alpha3 code
	fmt.Sprintf("country 3-letter code", geo.CountryAlpha2CodeToAlpha3Code(country.CountryAlpha2Code()))
}
```


