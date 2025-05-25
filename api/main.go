package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khareyash05/uptime-backend-api/cmd"
	"github.com/khareyash05/uptime-backend-api/types"
	db "github.com/khareyash05/uptime-backend-db"
	"github.com/khareyash05/uptime-backend-db/models"
	"github.com/rs/cors"

	"gorm.io/gorm"
)

var dbClient *gorm.DB

func init() {
	cmd.InitDB()
	dbClient = db.GetDB()
}

func main() {
	router := gin.Default()

	// Configure CORS
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	// Apply CORS middleware
	router.Use(func(c *gin.Context) {
		corsConfig.HandlerFunc(c.Writer, c.Request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.POST("/api/v1/website", authMiddleware(), func(c *gin.Context) {
		var request types.RequestUser
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		website := &models.Website{
			UserId:   request.UserId,
			URL:      request.URL,
		}

		result := dbClient.Create(website)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create website: " + result.Error.Error(),
			})
			return
		}
		c.JSON(201, gin.H{
			"message": website.ID,
		})
	})

	router.GET("/api/v1/website/status", authMiddleware(), func(c *gin.Context) {
		websiteId := c.Query("id")
		var request types.RequestUser2
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		if websiteId == "" {
			c.JSON(400, gin.H{
				"error": "Website ID is required",
			})
			return
		}
		var website models.Website
		result := dbClient.Model(models.Website{
			ID:       websiteId,
			UserId:   request.UserId,
			Disabled: false,
		}).First(&website)
		if result.Error != nil {
			c.JSON(404, gin.H{
				"error": "Website not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"data": website,
		})

	})

	router.GET("/api/v1/websites", authMiddleware(), func(c *gin.Context) {
		var request types.RequestUser2
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		var websites []models.Website
		result := dbClient.Where("user_id = ?", request.UserId).Find(&websites)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to fetch websites",
			})
			return
		}

		c.JSON(200, gin.H{
			"data": websites,
		})
	})

	router.DELETE("/api/v1/website/:id", authMiddleware(), func(c *gin.Context) {
		var request types.RequestUser2
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		websiteId := c.Param("id")
		if websiteId == "" {
			c.JSON(400, gin.H{
				"error": "Website ID is required",
			})
			return
		}

		var website models.Website
		result := dbClient.Model(models.Website{
			ID:       websiteId,
			UserId:   request.UserId,
			Disabled: false,
		}).First(&website)
		if result.Error != nil {
			c.JSON(404, gin.H{
				"error": "Website not found",
			})
			return
		}

		website.Disabled = true
		if err := dbClient.Save(&website).Error; err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to update website",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Website disabled successfully",
		})
	})
	router.Run()
}
