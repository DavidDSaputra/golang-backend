package main

import (
	"log"

	"gin-backend/config"
	"gin-backend/models"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("file .env tidak ditemukan, menggunakan environment variable sistem")
	}

	config.InitDatabase()

	products := []models.Product{
		{Name: "Nasi Goreng Spesial", Price: 25000, Category: "Makanan", Stock: 50, Description: "Nasi goreng dengan telur dan ayam", ImageURL: "https://picsum.photos/400", IsActive: true},
		{Name: "Sate Ayam 10 Tusuk", Price: 20000, Category: "Makanan", Stock: 100, Description: "Sate ayam dengan bumbu kacang", ImageURL: "https://picsum.photos/401", IsActive: true},
		{Name: "Es Teh Manis", Price: 8000, Category: "Minuman", Stock: 200, Description: "Es teh manis segar", ImageURL: "https://picsum.photos/402", IsActive: true},
		{Name: "Kopi Susu", Price: 15000, Category: "Minuman", Stock: 150, Description: "Kopi susu kekinian", ImageURL: "https://picsum.photos/403", IsActive: true},
		{Name: "Ayam Bakar", Price: 30000, Category: "Makanan", Stock: 30, Description: "Ayam bakar dengan sambal", ImageURL: "https://picsum.photos/404", IsActive: true},
	}

	inserted := 0
	for _, p := range products {
		var existing models.Product
		err := config.DB.Where("name = ?", p.Name).First(&existing).Error
		if err == nil {
			continue
		}
		if createErr := config.DB.Create(&p).Error; createErr != nil {
			log.Printf("gagal insert %s: %v", p.Name, createErr)
			continue
		}
		inserted++
	}

	log.Printf("seed selesai: %d produk baru ditambahkan", inserted)
}
