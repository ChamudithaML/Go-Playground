package dto

type Game struct {
	Id        string `json:"Id" bson:"id"`
	Name      string `json:"Name" bson:"name"`
	Genre     string `json:"Genre" bson:"genre"`
	CreatedAt string `json:"CreatedAt" bson:"createdat"`
}
