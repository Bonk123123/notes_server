package models

import (
	"github.com/Bonk123123/notes_server/internal/database"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
    Description string `json:"description"`
	Status uint `gorm:"default:0" json:"status"`
	UserID uint
	NoteElements []NoteElement
}

func (note *Note) Save() (*Note, error) {
    err := database.Db.Create(&note).Error
    if err != nil {
        return &Note{}, err
    }
    return note, nil
}

func Delete(note_id uint) error {
    err := database.Db.Delete(note_id).Error
    if err != nil {
        return err
    }
    return nil
}

func FindAllNotes() ([]Note, error) {
    var notes []Note
    err := database.Db.Find(&notes).Error
    if err != nil {
        return notes, err
    }
    return notes, nil
}

func FindNoteById(id uint) (Note, error) {
    var note Note
    err := database.Db.Where("ID=?", id).Find(&note).Preload("NoteElements").Error
    if err != nil {
        return Note{}, err
    }
    return note, nil
}