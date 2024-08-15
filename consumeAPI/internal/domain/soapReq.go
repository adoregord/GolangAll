package domain

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XmlnsSoap string   `xml:"xmlns:soap,attr"`
	XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
	Body      SOAPBody `xml:"soap:Body"`
}

type SOAPBody struct {
	XMLName  xml.Name `xml:"soap:Body"`
	Multiply Multiply `xml:"Multiply"`
}

type Multiply struct {
	XMLName xml.Name `xml:"Multiply"`
	Xmlns   string   `xml:"xmlns,attr"`
	IntA    int      `xml:"intA"`
	IntB    int      `xml:"intB"`
}
