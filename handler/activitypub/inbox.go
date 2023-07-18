package activitypub

import (
	"fmt"
	"net/http"
)

func PostInboxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(501)
	w.Write([]byte(fmt.Sprintf("{\"status\": 501, \"error\": \"NOT_IMPLEMENTED\", \"message\": \"%s\"}", "endpoint not implemented")))
}
