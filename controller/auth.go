package controller

import (
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JWTPrivateKey *rsa.PrivateKey

func Authentication(context *fiber.Ctx) error {

	if context.GetReqHeaders()["X-Api-Key"] != "123456" {
		return context.SendStatus(401)
	}

	claims := jwt.MapClaims{
		"name":  "default",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(JWTPrivateKey)
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.JSON(fiber.Map{"token": tokenString})
}
