package models

type Game struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Developer   string  `json:"developer"`
	Price       float64 `json:"price"`
}
