package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func CreateNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func UpdateNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func DeleteNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}