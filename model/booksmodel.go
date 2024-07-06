package model

import "time"

type Book struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title"`
	ISBN          string    `json:"isbn" gorm:"unique"`
	PublishedDate time.Time `json:"published_date" gorm:"type:date"`
	DelFlag       bool      `json:"del_flag" gorm:"default:false"`
	AuthorID      uint      `json:"author_id"`
}
