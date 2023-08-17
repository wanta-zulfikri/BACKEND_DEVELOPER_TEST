package middleware

import (
	"employe/config"
	"fmt"
	"log"
	"strings"
	"time" 
	

	"github.com/golang-jwt/jwt/v4"
) 

func CreateJWT(id uint, email, name string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"exp"     :        expirationTime.Unix(),
		"id"      :        id, 
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	signedToken, err := token.SignedString([]byte(config.JWTKey))
	if err != nil {
			log.Println(err, "jwt")
			return "", err 
			
	}
	return signedToken, nil
}

func ValidateJWT(authHeader string) (uint, error) {
		if authHeader == "" {
			return 0, fmt.Errorf("missing Authorization header")
		}
		tokenString := strings.Replace(authHeader, "Bearer","",1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.JWTKey), nil  
		})
		if err != nil {
				return 0, fmt.Errorf("invalid or expired token")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return 0, fmt.Errorf("invalid or expired token")
		}
		id, ok := claims["id"].(float64)
		if !ok {
			return 0, fmt.Errorf("invalid or expired token")
		}
		return uint(id), nil 
} 

func ValidateJWTUsername(authHeader string) (string, error) {
		if authHeader == "" {
			return "", fmt.Errorf("missing Authorization header")
		}
		tokenString := strings.Replace(authHeader, "Bearer", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.JWTKey), nil 
		})
		if err != nil {
			return "", fmt.Errorf("invalid or expired token")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
				return "", fmt.Errorf("invalid or expired token")
		} 
		username, ok := claims["username"].(string)
		if !ok {
				return "", fmt.Errorf("invalid or expired token")
		}
		return username, nil
} 

type Claims struct {
	ID 		     uint     `json:"id"`
	FirstName    string   `json:"firstname"`
	LastName     string   `json:"lastname"`
}

func (c *Claims) Valid() error {
		if c.ID == 0 || c.FirstName == "" || c.LastName == "" {
			return fmt.Errorf("invalid claims")
		}
		return nil 
}

func ValidateJWT2(authHeader string) (*Claims, error) {
		if authHeader == "" {
			return nil, fmt.Errorf("missing Authorization header")
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.JWTKey), nil
		})
		if err != nil {
			return nil, fmt.Errorf("invalid or expired token")
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
				return claims, nil 
		} else {
			return nil, fmt.Errorf("invalid or expired token")
		}
}