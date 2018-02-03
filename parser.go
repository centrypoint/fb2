// Package fb2 represent .fb2 format parser
package fb2

import (
	"bytes"
	"encoding/xml"
	"io"
)

// Parser struct
type Parser struct {
	book []byte
}

// New creates new Parser
func New(data []byte) *Parser {
	return &Parser{
		book: data,
	}
}

// CharsetReader required for change encodings
func (p *Parser) CharsetReader(c string, i io.Reader) (r io.Reader, e error) {
	switch c {
	case "windows-1251":
		r = decodeWin1251(i)
	}
	return
}

// Unmarshal parse data to FB2 type
func (p *Parser) Unmarshal() (result FB2, err error) {
	reader := bytes.NewReader(p.book)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = p.CharsetReader
	if err = decoder.Decode(&result); err != nil {
		return
	}

	result.UnmarshalCoverpage(p.book)

	return
}
