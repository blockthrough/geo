package geo

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/oschwald/geoip2-golang"
)

// Reader - a thin wrapper which provides helper function to parse information provided by embedded geoip2.Reader
type Reader struct {
	*geoip2.Reader
}

// NewMaxMindFromReader - create a new MaxMind Reader from io.Reader
func NewMaxMindFromReader(r io.Reader) (*Reader, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	return NewMaxMindFromBytes(data)
}

// NewMaxMindFromBytes - create a new MaxMind Reader from bytes
func NewMaxMindFromBytes(b []byte) (*Reader, error) {
	reader, err := geoip2.FromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("maxminddb.FromBytes: %w", err)
	}
	return &Reader{Reader: reader}, nil
}

// Close - close the reader
func (m *Reader) Close() error {
	return m.Reader.Close()
}

// Country - return an country object by passing a net.IP struct
func (m *Reader) Country(ip net.IP) (*country, error) {
	record, err := m.Reader.Country(ip)
	if err != nil {
		return nil, fmt.Errorf("fail to retrieve country with ip:%s, err:%w", ip, err)
	}

	c := country(*record)
	return &c, nil
}

// CountryByIPString - return an country object by passing an IPV4/IPV6 string
func (m *Reader) CountryByIPString(ip string) (*country, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return nil, fmt.Errorf("invalid ip: %s", ip)
	}

	return m.Country(parsedIP)
}

// IsTestDB - check if the underling DB is a test example, useful for sanity checks
func (m *Reader) IsTestDB() bool {
	desc, ok := m.Reader.Metadata().Description["en"]
	if !ok {
		return false
	}

	return strings.Contains(desc, "Test Database")
}

// BuildTimestamp - the timestamp when the MaxMind DB is built
func (m *Reader) BuildTimestamp() time.Time {
	return time.Unix(int64(m.Metadata().BuildEpoch), 0)
}

// Version - the opinionated semantic version for the underlying MaxMind DB in the format of v<major_version>.<minor_version>
func (m *Reader) Version() string {
	meta := m.Metadata()
	return fmt.Sprintf("v%d.%d", meta.BinaryFormatMajorVersion, meta.BinaryFormatMinorVersion)
}
