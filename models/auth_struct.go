package models

type EventBooking struct {
	ClientName   string `json:"clientName"`
	EventType    string `json:"eventType"`
	Date         string `json:"date"`
	MusicianType string `json:"musicianType"`
	OrderNumber  string `json:"orderNumber"`
}

type BookingsResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}
