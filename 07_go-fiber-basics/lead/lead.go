package lead

import (
	"github.com/Kiteretzu/golang-projects/07_go-fiber-basics/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int64 `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	var lead = new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&lead)

	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(503).Send("No lead found with ID")
		return
	}

	db.Delete(&lead)
	c.Send("Lead successfully Deleted")
}
