package apub

type Activity struct {
	Context []string `json:"@context,omitempty"`
	ID      string   `json:"id"`
	Type    string   `json:"type"`
	Actor   string   `json:"actor"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Object  any      `json:"object"`
}
