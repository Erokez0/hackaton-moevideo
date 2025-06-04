package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Erokez0/hackaton-moevideo/src/config"
)

type Category struct {
	ID     int    `gorm:"primary_key" json:"id,omitempty"`
	Name   string `gorm:"uniqueIndex" json:"name,omitempty"`
	Parent string `json:"parent,omitempty"`
}

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBConfig.DbHost, config.DBConfig.DbUser, config.DBConfig.DbPassword, config.DBConfig.DbName, config.DBConfig.DbPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.AutoMigrate(&Category{}); err != nil {
		log.Fatal(err)
	}
}

func FindAll() []Category {
	var categories []Category
	DB.Find(&categories)
	return categories
}
func FindIdsLikeName(name string, confident bool) []int {
	var categories []Category
	var ids = map[int]bool{}

	wordsOfName := strings.Split(name, " ")
	if confident {
		DB.Raw("SELECT id FROM categories WHERE name ILIKE ?",
			fmt.Sprintf("%%%s%%", name)).
			Scan(&categories)
	} else {
		for _, word := range wordsOfName {
			if len(word) < 3 {
				continue
			}
			DB.Raw("SELECT id FROM categories WHERE name ILIKE ? OR parent ILIKE ? AND @(length(parent) - length(?)) < length(?)",
				fmt.Sprintf("%%%s%%", word), fmt.Sprintf("%%%s%%", word), word, word).
				Scan(&categories)

		}
	}

	for _, category := range categories {
		if _, ok := ids[category.ID]; !ok {
			ids[category.ID] = true
		}
	}

	var result []int
	for id := range ids {
		result = append(result, id)
	}
	return result
}
func Seed() {
	var count int64
	DB.Model(&Category{}).Count(&count)
	if count > 0 {
		log.Println("\x1b[33mDatabase is already seeded\x1b[0m")
		return
	}

	data, err := os.ReadFile("categories.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var categories []Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	if DB.Create(&categories).Error != nil {
		log.Fatalf("Error creating categories: %v", err)
	}

	log.Println("\x1b[32mDatabase seeded\x1b[0m")
}
