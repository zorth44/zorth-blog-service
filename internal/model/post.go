package model

type Post struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Slug     string   `json:"slug" gorm:"uniqueIndex"`
	Title    string   `json:"title"`
	Date     string   `json:"date"`
	Category string   `json:"category"`
	Tags     []string `json:"tags" gorm:"serializer:json"`
	Excerpt  string   `json:"excerpt"`
}
