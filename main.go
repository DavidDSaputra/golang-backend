package main

import (
	"log"
	"os"

	"gin-backend/config"
	"gin-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("file .env tidak ditemukan, menggunakan environment variable sistem")
	}

	config.InitFirebase()
	config.InitDatabase()

	router := routes.SetupRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server berjalan di http://localhost:%s", port)
	log.Printf("health check: http://localhost:%s/v1/health", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("gagal menjalankan server: %v", err)
	}
}
