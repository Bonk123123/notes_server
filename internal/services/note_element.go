package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetNoteElements(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetNoteElement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func CreateNoteElement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func UpdateNoteElement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func DeleteNoteElement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}