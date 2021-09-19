package model

type Question struct {
	ID       int      `gorm:"primaryKey;unique;notNull;autoIncrement" json:"id"`
	Question string   `json:"question" binding:"required"`
	Answers  []Answer `json:"answers"`
}
