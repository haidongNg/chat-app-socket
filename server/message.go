package main

type Message struct {
	Message  string `json:"message"`
	ClientId string `json:"clientId"`
	RoomId   string `json:"roomId"`
}
