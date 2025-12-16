package model

// Chat struct in Go
type Chat struct {
	ID        string   `json:"id" bson:"_id,omitempty"`
	Users     []string `json:"users" db:"users" bson:"users"`
	CreatedAt string   `json:"createdAt" db:"createdAt" bson:"createdAt"`
}
