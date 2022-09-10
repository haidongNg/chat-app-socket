package main

type Room struct {
	RoomId   string             `json:"roomId"`
	RoomName string             `json:"roomName"`
	Clients  map[string]*Client `json:"clients"`
}
