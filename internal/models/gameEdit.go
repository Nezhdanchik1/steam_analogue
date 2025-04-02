package models

type GameEdit struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Developer   string  `json:"developer"`
	Price       float64 `json:"price"`
}
