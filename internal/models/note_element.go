package models

import (
	"github.com/Bonk123123/notes_server/internal/database"
	"gorm.io/gorm"
)

type NoteElement struct {
	gorm.Model
	Type string `gorm:"size:30;not null" json:"type"`
	Importance int `gorm:"default:0" json:"importance"`
    Content string `gorm:"not null;" json:"content"`
	NoteID uint
}

func (note_element *NoteElement) Save() (*NoteElement, error) {
    err := database.Db.Create(&note_element).Error
    if err != nil {
        return &NoteElement{}, err
    }
    return note_element, nil
}

func FindNoteElementById(id uint) (NoteElement, error) {
    var note_element NoteElement
    err := database.Db.Where("ID=?", id).Find(&note_element).Error
    if err != nil {
        return NoteElement{}, err
    }
    return note_element, nil
}