package model

type Author struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	DelFlag bool   `json:"del_flag" gorm:"default:false"`
}
