package main

import (
	"fmt"
	"github.com/103cuong/fiber_gorm_rest_api/cat"
	"github.com/103cuong/fiber_gorm_rest_api/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func initDatabase()  {
	var err error
	database.DB, err = gorm.Open("sqlite3", "cats.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")
	database.DB.AutoMigrate(&cat.Cat{})
	fmt.Println("database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	// TODO: move to router.go file
	app.Get("/api/cats", cat.GetCats)
	app.Get("/api/cats/:id", cat.GetCat)
	app.Post("/api/cats", cat.CreateCat)
	app.Put("/api/cats/:id", cat.UpdateCat)
	app.Delete("/api/cats/:id", cat.DeleteCat)

	app.Listen(3000)

	defer database.DB.Close()
}
