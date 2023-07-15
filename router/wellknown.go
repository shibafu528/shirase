package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/shibafu528/shirase/handler/wellknown"
)

func WellKnown(r chi.Router) {
	r.Route("/.well-known", func(r chi.Router) {
		r.Get("/webfinger", wellknown.WebFingerHandler)
	})
}
