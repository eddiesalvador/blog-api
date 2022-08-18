package models

type Tag struct {
	ID 			uint 	`gorm:"primary_key column:id" json:"id"`
	Title 		string	`gorm:"column:title" json:"title"`
	Slug 		string	`gorm:"column:slug" json:"slug"`
	Content 	string	`gorm:"column:content" json:"content"`
}

func (Tag) TableName() string {
	return "tag"
}