package models

import "time"

type Post struct {
	Id          string    `json:"id"`
	PostContent string    `json:"post_content"`
	CreateAt    time.Time `json:"create_at"`
	UserId      string    `json:"User_id"`
}
