package model

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Isbn  string `json:"isbn"`
	Price int    `json:"price"`
}
