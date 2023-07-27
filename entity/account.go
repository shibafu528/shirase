package entity

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/apub"
)

type Account struct {
	ID            int64
	Username      string
	Domain        sql.NullString
	DisplayName   sql.NullString
	PrivateKey    string
	PublicKey     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ActivityPubID sql.NullString
	Description   sql.NullString
}

func (a *Account) PreferredActivityPubID() string {
	if a.ActivityPubID.Valid {
		return a.ActivityPubID.String
	} else {
		return a.Username
	}
}

func (a *Account) ActorEndpointURL() *url.URL {
	return shirase.GlobalConfig.URLBase().JoinPath("users", a.PreferredActivityPubID())
}

func (a *Account) ActivityPubPerson() *apub.Person {
	actorEndpoint := a.ActorEndpointURL()
	return &apub.Person{
		Context:           []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"},
		ID:                actorEndpoint.String(),
		Type:              "Person",
		Inbox:             actorEndpoint.JoinPath("inbox").String(),
		Outbox:            actorEndpoint.JoinPath("outbox").String(),
		PreferredUsername: a.Username,
		Name:              a.DisplayName.String,
		Summary:           a.Description.String,
		PublicKey: apub.PublicKey{
			ID:           actorEndpoint.String() + "#main-key",
			Owner:        actorEndpoint.String(),
			PublicKeyPem: a.PublicKey,
		},
	}
}
