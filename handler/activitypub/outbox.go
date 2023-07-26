package activitypub

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shibafu528/shirase/apub"
	"github.com/shibafu528/shirase/db"
	"github.com/shibafu528/shirase/entity"
	"github.com/shibafu528/shirase/repo"
)

func GetOutboxHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("{\"status\": 404, \"error\": \"NOT_FOUND\", \"message\": \"%s\"}", "user not found")))
		return
	}

	_, q := db.Open()
	aid, err := q.GetAccountIDByActivityPubID(r.Context(), sql.NullString{String: username, Valid: true})
	if errors.Is(err, sql.ErrNoRows) {
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

	statuses, err := repo.NewStatusRepository().GetStatusesByAccountID(r.Context(), aid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("{\"status\": 500, \"error\": \"INTERNAL_ERROR\", \"message\": \"%s\"}", "internal error")))
		return
	}

	var activities []*apub.Activity
	for _, s := range statuses {
		a := s.ActivityPubActivity()
		a.Context = nil
		activities = append(activities, a)
	}

	actorEndpoint := entity.ActorEndpointByID(username)
	collection := apub.NewOrderedCollection(actorEndpoint.Outbox().String(), activities)

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(collection)
}
