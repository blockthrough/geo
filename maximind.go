package geo

import (
	"fmt"
	"io"
	"net"

	"github.com/oschwald/geoip2-golang"
)

type Reader struct {
	reader *geoip2.Reader
}

func NewMaxMindReader(r io.Reader) (*Reader, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	reader, err := geoip2.FromBytes(data)
	if err != nil {
		return nil, fmt.Errorf("maxminddb.FromBytes: %w", err)
	}

	return &Reader{reader: reader}, nil
}

func (m *Reader) Close() {
	m.reader.Close()
}

func (m *Reader) Country(ip net.IP) (*Country, error) {
	record, err := m.reader.Country(ip)
	if err != nil {
		return nil, fmt.Errorf("fail to retrieve country with ip:%s, err:%w", ip, err)
	}

	c := Country(*record)
	return &c, nil
}

func (m *Reader) CountryByIPString(ip string) (*Country, error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return nil, fmt.Errorf("invalid ip: %s", ip)
	}

	return m.Country(parsedIP)
}
