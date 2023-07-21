package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shibafu528/shirase/handler/activitypub"
	"github.com/shibafu528/shirase/handler/wellknown"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/.well-known", func(r chi.Router) {
		r.Get("/host-meta", wellknown.HostMetaHandler)
		r.Get("/webfinger", wellknown.WebFingerHandler)
	})
	r.Route("/users/{username}", func(r chi.Router) {
		r.Get("/", activitypub.GetPersonHandler)
		r.Post("/inbox", activitypub.PostInboxHandler)
		r.Get("/outbox", activitypub.GetOutboxHandler)
	})

	return r
}
