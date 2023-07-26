package apub

type OrderedCollection[T any] struct {
	Context      []string `json:"@context,omitempty"`
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	TotalItems   uint     `json:"totalItems"`
	OrderedItems []T      `json:"orderedItems,omitempty"`
}

func NewOrderedCollection[T any](id string, items []T) *OrderedCollection[T] {
	return &OrderedCollection[T]{
		Context:      []string{"https://www.w3.org/ns/activitystreams"},
		ID:           id,
		Type:         "OrderedCollection",
		TotalItems:   uint(len(items)),
		OrderedItems: items,
	}
}
