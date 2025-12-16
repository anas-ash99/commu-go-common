package model

type CallType string
type CallStatus string

const (
	Audio    CallType   = "audio"
	Video    CallType   = "video"
	Ringing  CallStatus = "ringing"
	Canceled CallStatus = "canceled"
	NoAnswer CallStatus = "no_answer"
	Ended    CallStatus = "ended"
	Missed   CallStatus = "missed"
	Declined CallStatus = "declined"
)

type Call struct {
	ID           string     `json:"id" bson:"_id"`
	Type         CallType   `json:"type" bson:"type"`
	Status       CallStatus `json:"status" bson:"status"`
	CallerID     string     `json:"callerId" bson:"callerId"`
	ReceiverId   string     `json:"receiverId" bson:"receiverId"`
	ReceiverName string     `json:"receiverName" bson:"receiverName"`
	CallerName   string     `json:"callerName" bson:"callerName"`
	StartedAt    *string    `json:"startedAt,omitempty" bson:"startedAt,omitempty"`
	EndedAt      *string    `json:"endedAt,omitempty" bson:"endedAt,omitempty"`
	CreatedAt    string     `json:"createdAt" bson:"createdAt"`
	Duration     *int       `json:"duration,omitempty" bson:"duration,omitempty"`
}

type CallResponse struct {
	CallId      string `json:"callId"`
	CallerId    string `json:"callerId"`
	ChannelName string `json:"channelName"`
	IsAccepted  string `json:"isAccepted"`
}

type TokenResponse struct {
	Token       string `json:"token"`
	ChannelName string `json:"channelName"`
}

type CallWithToken struct {
	Call         Call          `json:"call"`
	TokenRequest TokenResponse `json:"tokenRequest"`
}

type EndCallRequest struct {
	CallId    string `json:"CallId"`
	EndedAt   string `json:"endedAt"`
	StartedAt string `json:"startedAt"`
}

type UpdateCallRequest struct {
	CallId string     `json:"CallId"`
	Status CallStatus `json:"status"`
	UserId string     `json:"userId"` // user to be notified
}
