package model

type Url struct {
	Id       int
	Original string `json:"original"`
	ShortKey string `json:"shortKey"`
}
