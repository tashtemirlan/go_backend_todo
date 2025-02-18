package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"material_todo_go/config"
	"time"
)

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

// Parse JWT token and return email
func ParseJWT(tokenString string) (string, error) {
	fmt.Println("Received token:", tokenString) // Debugging output

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil // Ensure correct key type
	})

	if err != nil {
		fmt.Println("Error parsing token:", err) // Debugging output
		return "", errors.New("invalid token")
	}

	if !token.Valid {
		fmt.Println("Token is not valid") // Debugging output
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["email"] == nil {
		fmt.Println("Invalid claims structure") // Debugging output
		return "", errors.New("invalid claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		fmt.Println("Invalid email claim type") // Debugging output
		return "", errors.New("invalid email claim type")
	}

	return email, nil
}
