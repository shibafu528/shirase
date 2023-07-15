package db

import (
	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/apub"
)

func (a *Account) ActivityPubPerson() *apub.Person {
	actorEndpoint := shirase.GlobalConfig.URLBase().JoinPath("users", a.Username)
	return &apub.Person{
		Context:           []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"},
		ID:                actorEndpoint.String(),
		Type:              "Person",
		PreferredUsername: a.Username,
		Inbox:             actorEndpoint.JoinPath("inbox").String(),
		Outbox:            actorEndpoint.JoinPath("outbox").String(),
		PublicKey: apub.PublicKey{
			ID:           actorEndpoint.String() + "#main-key",
			Owner:        actorEndpoint.String(),
			PublicKeyPem: a.PublicKey,
		},
	}
}
