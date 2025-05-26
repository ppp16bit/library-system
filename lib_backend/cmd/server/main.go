package main

import (
	"log"
	"os"
	"time"

	"lib_backend/internal/config"
	handler "lib_backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadEnv()

	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("error configuring db: %v", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("error: %v", closeErr)
		}
	}()
	log.Println("sucess!")

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{
		"http://localhost:3000",
		"http://192.168.1.14:3000",
	}

	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	corsConfig.MaxAge = 12 * time.Hour

	r.Use(cors.New(corsConfig))

	handler.SetupRoutes(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s...", port)
	log.Fatal(r.Run(":" + port))
}
