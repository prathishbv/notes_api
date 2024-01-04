package main

// import (
// 	"fmt"

// 	"github.com/prathishbv/notes-api/model"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func main() {
// 	sqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
// 		"localhost",
// 		"postgres",
// 		"",
// 		"notes-app",
// 		"5432",
// 	)

// 	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect to the database")
// 	}
// 	fmt.Println("Connected successfully to the database")

// 	err = db.AutoMigrate(&model.Note{}, &model.SharedNote{})
// 	if err != nil {
// 		panic("failed to migrate database")
// 	}

// 	sqlDB, _ := db.DB()
// 	defer sqlDB.Close()
// }
