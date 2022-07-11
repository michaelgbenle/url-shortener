package models

type Short struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Shorted  string `json:"shorted" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
}
