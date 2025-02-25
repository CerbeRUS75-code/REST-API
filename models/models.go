package models

type Message struct {
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
