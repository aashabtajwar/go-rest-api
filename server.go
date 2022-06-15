package main

import "github.com/gin-gonic/gin"

func main() {
	// the server
	server := gin.Default()

	// Context is a struct
	// ctx is a pointer to that struct
	server.GET("/test", func(ctx *gin.Context) {

		// will return a JSON
		// first param is HTTP code
		ctx.JSON(200, gin.H{
			"message": "Okay",
		})
	})

	server.Run(":8000")
}
