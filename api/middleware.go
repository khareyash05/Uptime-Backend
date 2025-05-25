package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Printf("Authorization header: %q\n", token)
		if token == "" {
			fmt.Println("No Authorization header found, rejecting request")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		fmt.Println("Authorization header found, proceeding with request")
	}
}
