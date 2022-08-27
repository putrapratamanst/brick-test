package model

type DataProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageLink   string `json:"imageLink"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	Store       string `json:"store"`
}