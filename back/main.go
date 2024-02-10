package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

    // User routes
    r.POST("/register", Register)
    r.POST("/login", Login)
    r.GET("/getUsers", GetUsers)
    r.POST("/createConversation", CreateConversation)
    r.GET("/getConversations", GetConversations)
    r.DELETE("/deleteConversation", DeleteConversation)
    r.GET("/getConversationMessages", GetConversationMessages)
    r.POST("/createMessage", CreateMessage)

    r.Run(":8000") // Run on port 8000
}