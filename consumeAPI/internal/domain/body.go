package domain

import "encoding/xml"

type Body struct {
	XMLName          xml.Name         `xml:"Body"`
	MultiplyResponse MultiplyResponse `xml:"MultiplyResponse"`
}
