package main

import (
	"fmt"

	"github.com/Kiteretzu/golang-projects/07_go-fiber-basics/database"
	"github.com/Kiteretzu/golang-projects/07_go-fiber-basics/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
