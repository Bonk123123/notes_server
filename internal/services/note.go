package services

import (
	"net/http"
	"strconv"

	"github.com/Bonk123123/notes_server/internal/config"
	"github.com/Bonk123123/notes_server/internal/dto"
	"github.com/Bonk123123/notes_server/internal/models"
	"github.com/Bonk123123/notes_server/internal/pkg"
	"github.com/gin-gonic/gin"
)


func GetNotes(c *gin.Context) {

	notes, err := models.FindAllNotes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func GetNote(c *gin.Context) {
	note_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}


	note, err := models.FindNoteById(uint(note_id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

func CreateNote(c *gin.Context) {

	user, err := pkg.CurrentUser(c, config.REFRESH)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var input dto.NoteDto

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	note := models.Note{
		Name: input.Name,
   		Description: input.Description,
		Status: uint(config.IN_PROGRESS),
		UserID: user.ID,
	}

	note_save, err := note.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return	
	}

	c.JSON(http.StatusOK, note_save)
}

func UpdateNote(c *gin.Context) {
	var input dto.UpdateNoteDto

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	note_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}


	note, err := models.FindNoteById(uint(note_id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	note.Name = input.Name
	note.Description = input.Description
	note.Status = uint(input.Status)

	note_save, err := note.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return	
	}
	
	c.JSON(http.StatusOK, note_save)
}

func DeleteNote(c *gin.Context) {
	note_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	err = models.Delete(uint(note_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}