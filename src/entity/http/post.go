package http

import "time"

type PostRequest struct {
	Title 			string     	`json:"title"`
	Body 			string     	`json:"body"`
	UserId			uint		`json:"user_id"`
}

type PostDetailResponse struct {
	Id 				uint       	`json:"id"`
	Title 			string     	`json:"title"`
	Body		 	string     	`json:"body"`
	UserId			uint		`json:"user_id"`
	CreatedAt    	*time.Time 	`json:"created_at"`
	UpdatedAt    	*time.Time 	`json:"updated_at"`
}

type PostResponse struct {
	Id 				uint       	`json:"id"`
	Title 			string     	`json:"title"`
	CreatedAt    	*time.Time 	`json:"created_at"`
	UpdatedAt    	*time.Time 	`json:"updated_at"`
}