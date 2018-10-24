package model

// Dog defines a dog record
type Dog struct {
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Age         int    `json:"age"`
	Weight      string `json:"weight"`
	Condition   string `json:"condition"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Location    string `json:"location"`
	ImageURL    string `json:"image_url"`
}
