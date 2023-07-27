package entity

import (
	"net/url"
	"strconv"

	"github.com/shibafu528/shirase"
)

// Actor endpointのURLをラップし、いい感じのヘルパーを提供したりしなかったりする
type ActorEndpointHelper struct {
	*url.URL
	ActivityPubID string
}

type ActivityPubAccountIDProvider interface {
	ActivityPubAccountID() string
}

type ActivityPubAccountID string

func (i ActivityPubAccountID) ActivityPubAccountID() string {
	return string(i)
}

func ActorEndpoint(id ActivityPubAccountIDProvider) ActorEndpointHelper {
	apid := id.ActivityPubAccountID()
	return ActorEndpointHelper{shirase.GlobalConfig.URLBase().JoinPath("users", apid), apid}
}

func (a *ActorEndpointHelper) ActorPermalink() *url.URL {
	return shirase.GlobalConfig.URLBase().JoinPath("@" + a.ActivityPubID)
}

func (a *ActorEndpointHelper) StatusEndpoint(id int64) *url.URL {
	return a.JoinPath("statuses", strconv.FormatInt(id, 10))
}

func (a *ActorEndpointHelper) StatusPermalink(id int64) *url.URL {
	return a.ActorPermalink().JoinPath(strconv.FormatInt(id, 10))
}

func (a *ActorEndpointHelper) ActivityEndpoint(id int64) *url.URL {
	return a.JoinPath("statuses", strconv.FormatInt(id, 10), "activity")
}

func (a *ActorEndpointHelper) Outbox() *url.URL {
	return a.JoinPath("outbox")
}

func (a *ActorEndpointHelper) Following() *url.URL {
	return a.JoinPath("following")
}

func (a *ActorEndpointHelper) Followers() *url.URL {
	return a.JoinPath("followers")
}
