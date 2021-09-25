package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"linkbulb.io/links-service/models"
	"linkbulb.io/links-service/utils"
)

func setupLinksRoutes() {
	linksRouter := App.Group("/v1/links")
	linksRouter.Post("/new", createBulb)
}

func createBulb(c *fiber.Ctx) error {
	jwt := c.Cookies("session")
	if jwt == "" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Not logged in!",
		})
	}
	valid, userid := utils.ValidateJWT(jwt)
	if !valid {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "Invalid session token!",
		})
	}
	var bulbdto models.BulbDTO
	if err := c.BodyParser(&bulbdto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	uid := utils.RandStringBytes(8)
	links := bulbdto.LinksDTOToLinks()
	bulb := models.Bulb{
		UID:   uid,
		Type:  bulbdto.Type,
		User:  userid,
		Links: links,
	}
	models.DB.Create(bulb)
	return c.JSON(bulb)
}
