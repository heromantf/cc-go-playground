package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Welcome endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin Demo Server!",
		})
	})

	// API group with sample endpoints
	api := r.Group("/api/v1")
	{
		// Get all users
		api.GET("/users", func(c *gin.Context) {
			users := []gin.H{
				{"id": 1, "name": "Alice", "email": "alice@example.com"},
				{"id": 2, "name": "Bob", "email": "bob@example.com"},
				{"id": 3, "name": "Charlie", "email": "charlie@example.com"},
			}
			c.JSON(http.StatusOK, gin.H{
				"data": users,
			})
		})

		// Get user by ID
		api.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"data": gin.H{
					"id":    id,
					"name":  "Demo User",
					"email": "demo@example.com",
				},
			})
		})

		// Create a new user
		api.POST("/users", func(c *gin.Context) {
			var input struct {
				Name  string `json:"name" binding:"required"`
				Email string `json:"email" binding:"required,email"`
			}

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"message": "User created successfully",
				"data": gin.H{
					"name":  input.Name,
					"email": input.Email,
				},
			})
		})

		// Update user
		api.PUT("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			var input struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "User updated successfully",
				"data": gin.H{
					"id":    id,
					"name":  input.Name,
					"email": input.Email,
				},
			})
		})

		// Delete user
		api.DELETE("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
				"id":      id,
			})
		})
	}

	// Run server on port 8080
	r.Run(":8080")
}
