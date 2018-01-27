package fb2

import (
	"encoding/xml"
)

// List of interfaces for integration

// FB2 represents FB2 structure
type FB2 struct {
	FictionBook xml.Name   `xml:"FictionBook"`
	Stylesheet  []xml.Attr `xml:"stylesheet"`
	Description struct {
		TitleInfo struct {
			Genre      []string   `xml:"genre"`
			GenreType  []xml.Attr `xml:"genreType"`
			Author     []AuthorType
			BookTitle  string   `xml:"book-title"`
			Annotation string   `xml:"annotation"` // Check additional info
			Keywords   string   `xml:"keywords"`
			Date       string   `xml:"date"`
			Coverpage  struct { // Check additional info
				Image struct {
					Href string `xml:"xlink:href,attr"`
				} `xml:"image,allowempty"`
			} `xml:"coverpage"`
			Lang       string     `xml:"lang"`
			SrcLang    string     `xml:"src-lang"`
			Translator AuthorType `xml:"translator"` // AuthorType
			Sequence   string     `xml:"sequence"`   // SequenceType
		} `xml:"title-info"`
		DocumentInfo struct {
			Author      []AuthorType `xml:"author"` // AuthorType
			ProgramUsed string       `xml:"program-used"`
			Date        string       `xml:"date"`
			SrcURL      []string     `xml:"src-url"`
			SrcOcr      string       `xml:"src-ocr"`
			ID          string       `xml:"id"`
			Version     float64      `xml:"version"`
			History     string       `xml:"history"` // AnnotationType
		} `xml:"document-info"`
		PublishInfo struct {
			BookName  string `xml:"book-name"`
			Publisher string `xml:"publisher"`
			City      string `xml:"city"`
			Year      int    `xml:"year"`
			ISBN      string `xml:"isbn"`
			Sequence  string `xml:"sequence"` // SequenceType
		} `xml:"PublishInfo"`
		CustomInfo []struct {
			InfoType xml.Attr `xml:"info-type"`
		} `xml:"custom-info"`
	} `xml:"description"`
	Body struct {
		Sections []struct {
			P []string `xml:"p"`
		} `xml:"section"`
	} `xml:"body"`
	Binary []struct {
		ContentType xml.Attr `xml:"content-type"`
		ID          xml.Attr `xml:"id"`
	} `xml:"binary"`
}

// UnmarshalCoverpage func
func (f *FB2) UnmarshalCoverpage(data []byte) {
	tagOpened := false
	coverpageStartIndex := 0
	coverpageEndIndex := 0
	// imageHref := ""
	tagName := ""
_loop:
	for i, v := range data {
		if tagOpened {
			switch v {
			case '>':
				if tagName != "p" && tagName != "/p" {
				}
				tagOpened = false
				if tagName == "coverpage" {
					coverpageStartIndex = i + 1
				} else if tagName == "/coverpage" {
					coverpageEndIndex = i - 11
					break _loop
				}
				tagName = ""
				break
			default:
				tagName += string(v)
			}
		} else {
			if v == '<' {
				tagOpened = true
			}
		}
	}

	if coverpageEndIndex > coverpageStartIndex {
		href := parseImage(data[coverpageStartIndex:coverpageEndIndex])
		f.Description.TitleInfo.Coverpage.Image.Href = href
	}
}

// AuthorType embedded fb2 type, represents author info
type AuthorType struct {
	FirstName  string `xml:"first-name"`
	MiddleName string `xml:"middle-name"`
	LastName   string `xml:"last-name"`
	Nickname   string `xml:"nickname"`
	HomePage   string `xml:"home-page"`
	Email      string `xml:"email"`
}

// TextFieldType embedded fb2 type, represents text field
type TextFieldType struct {
}

// TitleType embedded fb2 type, represents title type fields
type TitleType struct {
	P         []string `xml:"p"`
	EmptyLine []string `xml:"empty-line"`
}

// ImageType embedded fb2 type, represents image information
type ImageType struct {
	Type xml.Attr `xml:"xlink:type"`
	Href xml.Attr `xml:"xlink:href"`
	Alt  string   `xml:"alt"`
}

// PType embedded fb2 type, represents paragraph
type PType struct {
}
