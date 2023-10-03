# Geo Information Library For Go
[![CircleCI](https://circleci.com/gh/blockthrough/geo.svg?style=svg&circle-token=b0554d26f90621f9996755fe9fd6665e74cabcbe)](<https://app.circleci.com/pipelines/github/blockthrough/geo?branch=maximind-geo>)

[![codecov](https://codecov.io/gh/blockthrough/geo/branch/maximind-geo/graph/badge.svg?token=DK6KZBFHML)](https://codecov.io/gh/blockthrough/geo)

The library currently supports reads MaxMind GeoIP2 database and provide a thin decorator/helper function on the returned data for easy retrival and deriving on top of the geo info provided by MaxMind. It also provides a normalized way to represent the geo info in standard ISO format including unknown geo info.

The implemenation of this library is based upon other 2 open source repositories:

1. https://github.com/biter777/countries 
2. https://github.com/oschwald/geoip2-golang

## Goal
The library is visioned to support more features related to processing of the geo info data in the future. That means, support reading from multiple geo database providers and expose more helpers on the data for commonly encountered user scenarios.


## Features
1. Support easy look up for geo information from MaxMind DB using ip string.
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


```


