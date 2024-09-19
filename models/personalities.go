package models

type Personality struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

var Personalities = []Personality{
	{Id: 1, Name: "John Doe", Bio: "A regular guy"},
	{Id: 2, Name: "Jane Doe", Bio: "A regular woman"},
}
