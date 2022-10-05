package models

type Alert struct {
	ID        string `json:"id" gorm:"primaryKey"`
	AlertName string `json:"alertName"`
	EventId   string `json:"eventId"`
	Host      string `json:"host"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Occurred  string `json:"occurred"`
	Origin    string `json:"origin"`
}
