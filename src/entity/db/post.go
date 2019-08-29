package db

import "time"

type Post struct {
	Id 				uint       	`gorm:"primary_key" json:"id"`
	Title 			string     	`gorm:"column:title" json:"title"`
	Body 			string     	`gorm:"column:body" json:"body"`
	CreatedAt    	*time.Time 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt    	*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    	*time.Time 	`gorm:"column:deleted_at" json:"deleted_at"`
	User	 		*User 		`gorm:"auto_preload, foreignkey:UserId, association_foreignkey:ID"`
	UserId 			uint 		`gorm:"column:user_id"`
}

func (Post) TableName() string {
	return "posts"
}
