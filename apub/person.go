package apub

type Person struct {
	Context           []string `json:"@context,omitempty"`
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	Inbox             string   `json:"inbox"`
	Outbox            string   `json:"outbox"`
	PreferredUsername string   `json:"preferredUsername"`
	Name              string   `json:"name"`
	Summary           string   `json:"summary,omitempty"`
	//Discoverable      bool      `json:"discoverable"`
	PublicKey PublicKey `json:"publicKey"`
}

type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}
