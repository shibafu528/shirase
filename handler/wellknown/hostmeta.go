package wellknown

import (
	"encoding/xml"
	"net/http"

	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/apub/xrd"
)

func HostMetaHandler(w http.ResponseWriter, r *http.Request) {
	m := xrd.NewHostMeta(shirase.GlobalConfig.URLBase())
	w.Header().Set("Content-Type", "application/xrd+xml; charset=utf-8")
	w.Write([]byte(xml.Header))
	xml.NewEncoder(w).Encode(m)
}
