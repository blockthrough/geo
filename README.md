# Geo Information Library For Go
[![CircleCI](https://circleci.com/gh/blockthrough/geo.svg?style=svg&circle-token=b0554d26f90621f9996755fe9fd6665e74cabcbe)](<https://app.circleci.com/pipelines/github/blockthrough/geo?branch=maximind-geo>)

[![codecov](https://codecov.io/gh/blockthrough/geo/branch/maximind-geo/graph/badge.svg?token=DK6KZBFHML)](https://codecov.io/gh/blockthrough/geo)

The library currently supports reads MaxMind GeoIP2 database and provide a thin decorator/helper function on the returned data for easy retrival and deriving on top of the geo info provided by MaxMind. It also provides a normalized way to represent the geo info in standard ISO format including the scenario of unknown.

The implemenation of this library is based upon other 2 open source repositories:

1. https://github.com/biter777/countries 
2. https://github.com/oschwald/geoip2-golang

## Goal
The library is visioned to support more features related to processing of the geo info data in the future. That means, support reading from multiple geo database providers and expose more helpers on the data for commonly encountered user scenarios.


## Features
1. Support easy lookup for geo information from MaxMind DB using ip string.
2. Support both alpha2 country code and alpha2 country code.
3. Support any `io.Reader` for user to provide the MaxMind mmdb file.
4. Provide default return value of `ZZ` and `ZZZ` as specified in ISO standard if country is unknown.


## Installation
```
GOPRIVATE=github.com/blockthrough get github.com/blockthrough/geo
```

Since it is not an open-sourced repository yet, use `GOPRIVATE` to make sure go installer knows it is a private module. You also need to make sure you have correct git access to the private module.


## How To Use

### Read Maxmind DB and LookUp Geo With IP
```go

package main 

import (
    "embed"
	"fmt"
)

////go:embed <your_maxmind_db>
var embedFS embed.FS

func main() {

    file, err := embedFS.Open(maxmindCountryTestDB)
	if err != nil {
		t.Fatal(fmt.Errorf("embedFS.Open: %w", err))
	}

	maximind, err := NewMaxMindReader(file)
	if err != nil {
		t.Fatal(fmt.Errorf("NewMaxMindn: %w", err))
	}

    // check if it is test DB
    fmt.Sprintf("test db?: %t",maxmind.IsUsingTestDB())

    // you can pass ipv4 or ipv6 
    country, err := maxmind.Country("127.0.0.1")
    if err != nil {
        fmt.Error("err:%s",err)
    }

    fmt.Sprintf("unknown country: %s", country.isUnknown())) // is he country unknown?
    fmt.Sprintf("country code: %s", country.CountryAlpha2Code()) // 2-letter country code
    fmt.Sprintf("country 3-letter code: %s", country.CountryAlpha3Code()) // 3-letter country code
    fmt.Sprintf("continent code: %s", country.ContinentCode()) // contintue code


    // you can also use adapter directly if we want to have alpha3 code
    fmt.Sprintf("country 3-letter code", CountryAlpha2CodeToAlpha3Code(country.CountryAlpha2Code()))
}


```


