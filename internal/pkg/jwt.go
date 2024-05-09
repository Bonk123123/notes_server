package pkg

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Bonk123123/notes_server/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type token_type string

const (
	ACCESS  token_type = "ACCESS_TOKEN_TTL"
	REFRESH token_type = "REFRESH_TOKEN_TTL"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user models.User, tt token_type) (string, error) {
    tokenTTL, _ := strconv.Atoi(os.Getenv(string(tt)))
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  user.ID,
        "iat": time.Now().Unix(),
        "eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
    })
    return token.SignedString(privateKey)
}

func ValidateJWT(c *gin.Context, tt token_type) error {
    token, err := getToken(c, tt)
    if err != nil {
        return err
    }
    _, ok := token.Claims.(jwt.MapClaims)
    if ok && token.Valid {
        return nil
    }
    return errors.New("invalid token provided")
}

func CurrentUser(c *gin.Context, tt token_type) (models.User, error) {
    err := ValidateJWT(c, tt)
    if err != nil {
        return models.User{}, err
    }
    token, _ := getToken(c, tt)
    claims, _ := token.Claims.(jwt.MapClaims)
    userId := uint(claims["id"].(float64))

    user, err := models.FindUserById(userId)
    if err != nil {
        return models.User{}, err
    }
    return user, nil
}

func getToken(c *gin.Context, tt token_type) (*jwt.Token, error) {
    var tokenString string
	if (tt == ACCESS) {
		tokenString = getTokenFromRequest(c)
	} else {
		tokenString = getTokenFromCookie(c)
	}
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return privateKey, nil
    })
    return token, err
}

func getTokenFromRequest(c *gin.Context) string {
    bearerToken := c.Request.Header.Get("Authorization")
    splitToken := strings.Split(bearerToken, " ")
    if len(splitToken) == 2 {
        return splitToken[1]
    }
    return ""
}

func getTokenFromCookie(c *gin.Context) string {
    bearerToken, err := c.Cookie("refresh_token")
	if err != nil {
		return ""
	}
    splitToken := strings.Split(bearerToken, " ")
    if len(splitToken) == 2 {
        return splitToken[1]
    }
    return ""
}