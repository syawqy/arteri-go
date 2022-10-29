package entities

type Record struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  uint   `json:"caregoryId"`
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
