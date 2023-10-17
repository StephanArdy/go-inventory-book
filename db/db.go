package db

import (
	"book-inventory/models"
	"log"
	"os"

	_ "database/sql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	return db

}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{}) //generate table

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		seederBook(db) //memasukkan data awal
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{{
		Title:       "Perjalanan ini",
		Author:      "Jojo",
		Description: "Buku tentang perjalanan",
		Stock:       10,
	}, {
		Title:       "Pengobata Herbal",
		Author:      "Sindy",
		Description: "Buku tentang pengobatan",
		Stock:       5,
	}, {
		Title:       "Seputar Hewan",
		Author:      "kiki",
		Description: "Buku tentang hewan",
		Stock:       7,
	}, {
		Title:       "Berita Terkini",
		Author:      "Sindy",
		Description: "Buku tentang berita",
		Stock:       5,
	}}

	for _, v := range data {
		db.Create(&v)
	}
}
