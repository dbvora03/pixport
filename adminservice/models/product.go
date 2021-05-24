package productmodel

type Product struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Dimensions  string `json:"dimensions"`
	Price       int    `json:"price"`
	Image       string `json:"image"` 
}