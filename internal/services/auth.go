package services

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Bonk123123/notes_server/internal/dto"
	"github.com/Bonk123123/notes_server/internal/models"
	"github.com/Bonk123123/notes_server/internal/pkg"
	"github.com/gin-gonic/gin"
)


func Auth(c *gin.Context)  {
	var input dto.AuthDto

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	
	user, err := models.FindUserByUsername(input.Username)

	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "incorect username or password"})
        return
    }

	err = user.ValidatePassword(input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorect username or password"})
	}

	access_jwt, a_err := pkg.GenerateJWT(user, pkg.ACCESS)
	refresh_jwt, r_err := pkg.GenerateJWT(user, pkg.REFRESH)

	if a_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": a_err.Error()})
        return
    }

	if r_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": r_err.Error()})
        return
    }

	ttl, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))

	c.SetCookie("refresh_token", string(refresh_jwt), ttl, "/", os.Getenv("HOST"), false, true)

	c.JSON(http.StatusOK, gin.H{"access_token": access_jwt})
}

func Registration(c *gin.Context)  {
	var input dto.RegistrationDto

	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if input.Password != input.PasswordRepeat {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords are different"})
	}

	user := models.User{
        Username: input.Username,
        Password: input.Password,
    }

	user.Save()

	access_jwt, a_err := pkg.GenerateJWT(user, pkg.ACCESS)
	refresh_jwt, r_err := pkg.GenerateJWT(user, pkg.REFRESH)

	if a_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": a_err.Error()})
        return
    }

	if r_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": r_err.Error()})
        return
    }

	ttl, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))

	c.SetCookie("refresh_token", string(refresh_jwt), ttl, "/", os.Getenv("HOST"), false, true)	

	c.JSON(http.StatusOK, gin.H{"access_token": access_jwt})
}

func Refresh(c *gin.Context)  {
	user, err := pkg.CurrentUser(c, pkg.REFRESH)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}


	access_jwt, a_err := pkg.GenerateJWT(user, pkg.ACCESS)
	refresh_jwt, r_err := pkg.GenerateJWT(user, pkg.REFRESH)

	if a_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": a_err.Error()})
        return
    }

	if r_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": r_err.Error()})
        return
    }

	ttl, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))

	c.SetCookie("refresh_token", string(refresh_jwt), ttl, "/", os.Getenv("HOST"), false, true)
	
	c.JSON(http.StatusOK, gin.H{"access_token": access_jwt})
}