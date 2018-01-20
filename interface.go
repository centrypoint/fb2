package fb2

import (
	"encoding/xml"
)

// List of interfaces for integration

// FB2 represents FB2 structure
type FB2 struct {
	FictionBook xml.Name `xml:"FictionBook"`
	Description struct {
		TitleInfo struct {
			Genre  string `xml:"genre"`
			Author struct {
				FirstName string `xml:"first-name"`
				LasName   string `xml:"last-name"`
			} `xml:"author"`
		} `xml:"title-info"`
	} `xml:"description"`
}
