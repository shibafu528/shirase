package apub

import "time"

type Note struct {
	Context      []string  `json:"@context,omitempty"`
	ID           string    `json:"id"`
	Type         string    `json:"type"`
	To           []string  `json:"to"`
	Cc           []string  `json:"cc"`
	Published    time.Time `json:"published"`
	URL          string    `json:"url"`
	AttributedTo string    `json:"attributedTo"`
	Content      string    `json:"content"`
}
