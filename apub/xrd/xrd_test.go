package xrd

import (
	"bytes"
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHostMeta(t *testing.T) {
	u, err := url.Parse("http://localhost")
	if err != nil {
		panic(err)
	}

	x := NewHostMeta(u)
	assert.Equal(t, &XRD{
		Xmlns: "http://docs.oasis-open.org/ns/xri/xrd-1.0",
		Links: []Link{
			{
				Rel:      "lrdd",
				Type:     "application/xrd+xml",
				Template: "http://localhost/.well-known/webfinger?resource={uri}",
			},
		},
	}, x)

	buf := &bytes.Buffer{}
	buf.Write([]byte(xml.Header))
	err = xml.NewEncoder(buf).Encode(x)
	assert.NoError(t, err)

	assert.Equal(t, `<?xml version="1.0" encoding="UTF-8"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0"><Link rel="lrdd" type="application/xrd+xml" template="http://localhost/.well-known/webfinger?resource={uri}"></Link></XRD>`, buf.String())
}
