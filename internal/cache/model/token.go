package model

import "encoding/json"

// FirebaseMessagingToken holds data related to Firebase high priority messaging token.
// This token is used when pushing notifications to clients.
type FirebaseMessagingToken struct {
	Token       string `json:"token"`
	SuperheroID string `json:"superheroID"`
	CreatedAt   string `json:"created_at"`
}

// MarshalBinary ...
func (f FirebaseMessagingToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(f)
}

// UnmarshalBinary ...
func (f *FirebaseMessagingToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &f)
}