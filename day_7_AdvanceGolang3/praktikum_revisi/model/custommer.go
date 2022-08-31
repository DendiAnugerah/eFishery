package model

type Custommer struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"string"`
}

type Custommers []Custommer
