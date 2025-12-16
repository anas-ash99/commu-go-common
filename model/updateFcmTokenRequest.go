package model

type UpdateFcmTokenRequest struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}
