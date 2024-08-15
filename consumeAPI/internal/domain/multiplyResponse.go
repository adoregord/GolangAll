package domain

import "encoding/xml"

type MultiplyResponse struct {
	XMLName        xml.Name `xml:"MultiplyResponse"`
	MultiplyResult string   `xml:"MultiplyResult"`
}
