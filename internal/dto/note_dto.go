package dto

import "github.com/Bonk123123/notes_server/internal/config"

type NoteDto struct {
	Name string `json:"name" binding:"required"`
    Description string `json:"description" binding:"required"`
}


type UpdateNoteDto struct {
	Name string `json:"name" binding:"required"`
    Description string `json:"description" binding:"required"`
    Status config.Status `json:"status" binding:"required"`
}