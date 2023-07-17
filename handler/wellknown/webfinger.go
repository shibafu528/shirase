package wellknown

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/apub"
	"github.com/shibafu528/shirase/db"
)

type WebFingerResponse struct {
	Subject string          `json:"subject"`
	Links   []WebFingerLink `json:"links"`
}

type WebFingerLink struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

func WebFingerHandler(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Query().Get("resource")
	if res == "" {
		w.WriteHeader(400)
		return
	}

	// TODO: impl find by url pattern
	acct, err := apub.ParseFullAcct(res)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("{\"status\": 422, \"error\": \"INVALID_RESOURCE_FORMAT\", \"message\": \"%s\"}", err)))
		return
	}
	if acct.Domain() != shirase.GlobalConfig.LocalDomain {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("{\"status\": 404, \"error\": \"NOT_FOUND\", \"message\": \"%s\"}", "resource not found")))
		return
	}

	_, q := db.Open()
	a, err := q.GetAccountByUsername(r.Context(), acct.Username())
	if errors.Is(err, sql.ErrNoRows) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("{\"status\": 404, \"error\": \"NOT_FOUND\", \"message\": \"%s\"}", "resource not found")))
		return
	} else if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("{\"status\": 500, \"error\": \"INTERNAL_ERROR\", \"message\": \"%s\"}", "internal error")))
		return
	}

	resp := WebFingerResponse{
		Subject: "acct:" + a.Username + "@" + shirase.GlobalConfig.LocalDomain,
		Links: []WebFingerLink{
			{Rel: "self", Type: "application/activity+json", Href: a.ActorEndpointURL().String()},
		},
	}

	w.Header().Set("Content-Type", "application/jrd+json")
	json.NewEncoder(w).Encode(resp)
}
