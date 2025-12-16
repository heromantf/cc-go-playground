package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for demo purposes
	},
}

// handleWebSocket handles WebSocket connections with an echo server
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket client connected")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		log.Printf("Received message: %s", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Failed to write message: %v", err)
			break
		}
	}

	log.Println("WebSocket client disconnected")
}

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

	// WebSocket endpoint - echoes messages back to the client
	r.GET("/ws", handleWebSocket)

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
