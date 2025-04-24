package middleware

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// JWT Secret loaded from environment variable
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// AuthMiddleware will check for a valid JWT token in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if it's present
		if len(tokenString) > 7 && strings.ToLower(tokenString[:7]) == "bearer " {
			tokenString = tokenString[7:]
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check if the token's method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token signature method")
			}
			
			
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract the student_id from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		studentID, ok := claims["student_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Student ID missing from token"})
			c.Abort()
			return
		}

		// Set the student ID in the context for use in subsequent handlers
		c.Set("student_id", uint(studentID))

		// Proceed to the next middleware or handler
		c.Next()
	}
}
