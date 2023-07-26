package entity

import (
	"net/url"
	"strconv"

	"github.com/shibafu528/shirase"
)

// Actor endpointのURLをラップし、いい感じのヘルパーを提供したりしなかったりする
type ActorEndpoint struct {
	*url.URL
	ActivityPubID string
}

func ActorEndpointByID(apid string) ActorEndpoint {
	return ActorEndpoint{shirase.GlobalConfig.URLBase().JoinPath("users", apid), apid}
}

func (a *ActorEndpoint) ActorPermalink() *url.URL {
	return shirase.GlobalConfig.URLBase().JoinPath("@" + a.ActivityPubID)
}

func (a *ActorEndpoint) StatusEndpoint(id int64) *url.URL {
	return a.JoinPath("statuses", strconv.FormatInt(id, 10))
}

func (a *ActorEndpoint) StatusPermalink(id int64) *url.URL {
	return a.ActorPermalink().JoinPath(strconv.FormatInt(id, 10))
}

func (a *ActorEndpoint) ActivityEndpoint(id int64) *url.URL {
	return a.JoinPath("statuses", strconv.FormatInt(id, 10), "activity")
}

func (a *ActorEndpoint) Outbox() *url.URL {
	return a.JoinPath("outbox")
}

func (a *ActorEndpoint) Following() *url.URL {
	return a.JoinPath("following")
}

func (a *ActorEndpoint) Followers() *url.URL {
	return a.JoinPath("followers")
}
