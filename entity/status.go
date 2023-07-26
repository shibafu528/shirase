package entity

import (
	"time"

	"github.com/shibafu528/shirase/apub"
)

type Status struct {
	ID            int64
	AccountID     int64
	ActivityPubID string
	Text          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s *Status) ActivityPubActivity() *apub.Activity {
	actorEndpoint := ActorEndpointByID(s.ActivityPubID)
	note := s.ActivityPubNote()
	note.Context = nil
	return &apub.Activity{
		Context: []string{"https://www.w3.org/ns/activitystreams"},
		ID:      actorEndpoint.ActivityEndpoint(s.ID).String(),
		Type:    "Create",
		Actor:   actorEndpoint.String(),
		To:      note.To,
		Cc:      note.Cc,
		Object:  note,
	}
}

func (s *Status) ActivityPubNote() *apub.Note {
	actorEndpoint := ActorEndpointByID(s.ActivityPubID)
	return &apub.Note{
		Context:      []string{"https://www.w3.org/ns/activitystreams"},
		ID:           actorEndpoint.StatusEndpoint(s.ID).String(),
		Type:         "Note",
		To:           []string{apub.PublicCollectionURI},
		Cc:           []string{actorEndpoint.Followers().String()},
		Published:    s.CreatedAt,
		URL:          actorEndpoint.StatusPermalink(s.ID).String(),
		AttributedTo: actorEndpoint.String(),
		Content:      s.Text,
	}
}
