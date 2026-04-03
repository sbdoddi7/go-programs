package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/slow", func(ctx *gin.Context) {
		log.Println("Started /slow request")
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{"message": "slow request"})
		log.Println("Finished /slow request")
	})

	log.Println("Starting the server :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Listen error :%s", err)
	}
}
