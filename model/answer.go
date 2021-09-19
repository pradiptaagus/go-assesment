package model

type Answer struct {
	ID         int `gorm:"primaryKey;unique;notNull;autoIncrement"`
	Name       string
	QuestionID int
	Question   Question
}
