package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/shibafu528/shirase/handler/activitypub"
)

func ActivityPub(r chi.Router) {
	r.Route("/users/{username}", func(r chi.Router) {
		r.Get("/", activitypub.GetPersonHandler)
		r.Post("/inbox", activitypub.PostInboxHandler)
		r.Get("/outbox", activitypub.GetOutboxHandler)
	})
}
