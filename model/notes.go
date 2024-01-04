package model

import "time"

type Note struct {
	ID         int       `gorm:"type:int;primary_key"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"type:timestamp without time zone;not null"`
	UpdatedAt  time.Time `gorm:"type:timestamp without time zone;not null"`
	CreatedByID int      `gorm:"type:int;not null"`
}

type SharedNote struct {
	ID        int       `gorm:"type:int;primary_key"`
	NoteID    int       `gorm:"type:int;not null"`
	UserID    int       `gorm:"type:int;not null"`
	SharedByID int      `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone;not null"`
}
