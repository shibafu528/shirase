package activitypub

import (
	"fmt"
	"net/http"
)

func PostInboxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(405)
	w.Write([]byte(fmt.Sprintf("{\"status\": 405, \"error\": \"NOT_IMPLEMENTED\", \"message\": \"%s\"}", "endpoint not implemented")))
}
