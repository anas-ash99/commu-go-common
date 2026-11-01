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

type OfferRequest struct {
	TargetUserID string `json:"targetUserId"`
	SDP          string `json:"sdp"`
}

type AnswerRequest struct {
	TargetUserID string `json:"targetUserId"`
	SDP          string `json:"sdp"`
}

type IceCandidateRequest struct {
	TargetUserID  string `json:"targetUserId"`
	Candidate     string `json:"candidate"`
	SdpMid        string `json:"sdpMid"`
	SdpMLineIndex int    `json:"sdpMLineIndex"`
}

type EndCallRequest struct {
	TargetUserID string `json:"targetUserId"`
	Reason       string `json:"reason"` // "hangup", "rejected", "busy"
}
