package cat

import (
	"github.com/103cuong/fiber_gorm_rest_api/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Cat struct {
	gorm.Model
	Name string `json:"name"`
	Color string `json:"color"`
	Age uint `json:"age"`
}

func GetCats(c *fiber.Ctx)  {
	db := database.DB
	var cats []Cat
	db.Find(&cats)
	c.JSON(cats)
}

func GetCat(c *fiber.Ctx)  {
	id := c.Params("id")
	db := database.DB
	var cat Cat
	db.Find(&cat, id)
	c.JSON(cat)
}

func CreateCat(c *fiber.Ctx)  {
	db := database.DB
	cat := new(Cat)
	if err := c.BodyParser(cat); err != nil {
		c.Status(503).JSON(err)
		return
	}
	db.Create(&cat)
	c.JSON(cat)
}

func UpdateCat(c *fiber.Ctx)  {
	id := c.Params("id")
	db := database.DB
	cat := new(Cat)
	db.Find(&cat, id)
	if cat.Name == "" {
		c.Status(500).JSON("no cat found with id")
	}
	if err := c.BodyParser(cat); err != nil {
		c.Status(503).JSON(err)
		return
	}
	db.Save(cat)
	c.JSON(cat)
}

func DeleteCat(c *fiber.Ctx)  {
	id := c.Params("id")
	db := database.DB
	var cat Cat
	db.First(&cat, id)
	if cat.Name == "" {
		c.Status(500).JSON("no cat found with id")
	}
	db.Delete(&cat)
	c.JSON("cat deleted")
}
