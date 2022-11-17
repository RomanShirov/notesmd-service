package crypto

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func GenerateUserToken(uid int) string {
	claims := jwt.MapClaims{
		"user_id": uid,
		"exp":     time.Now().Add(time.Hour * 8766).Unix(),
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return token
}

func GetUserIdFromToken(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(int)
	return userId
}
