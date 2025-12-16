package model

// User struct in Go
type User struct {
	ID          string   `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Friends     []string `json:"friends" bson:"friends"`
	IsOnline    bool     `json:"isOnline" bson:"isOnline"`
	IsTyping    bool     `json:"isTyping" bson:"isTyping"`
	FcmToken    string   `json:"fcmToken" bson:"fcmToken"`
	PhoneNumber string   `json:"phoneNumber"  bson:"phoneNumber"`
	CreatedAt   string   `json:"createdAt"  bson:"createdAt"`
}
