package activitypub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shibafu528/shirase/repo"
)

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("{\"status\": 404, \"error\": \"NOT_FOUND\", \"message\": \"%s\"}", "user not found")))
		return
	}

	a, err := repo.NewAccountRepository().GetAccountByActivityPubID(r.Context(), username)
	if errors.Is(err, repo.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("{\"status\": 404, \"error\": \"NOT_FOUND\", \"message\": \"%s\"}", "user not found")))
		return
	} else if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("{\"status\": 500, \"error\": \"INTERNAL_ERROR\", \"message\": \"%s\"}", "internal error")))
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(a.ActivityPubPerson())
}
