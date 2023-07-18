package xrd

import (
	"encoding/xml"
	"net/url"
)

type XRD struct {
	XMLName xml.Name `xml:"XRD"`
	Xmlns   string   `xml:"xmlns,attr"`
	Links   []Link
}

type Link struct {
	XMLName  xml.Name `xml:"Link"`
	Rel      string   `xml:"rel,attr"`
	Template string   `xml:"template,attr"`
}

func NewHostMeta(base *url.URL) *XRD {
	return &XRD{
		Xmlns: "http://docs.oasis-open.org/ns/xri/xrd-1.0",
		Links: []Link{
			{
				Rel:      "lrdd",
				Template: base.JoinPath(".well-known/webfinger").String() + "?resource={uri}",
			},
		},
	}
}
