package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type InventoryController interface {
	Route(app *fiber.App)
}

type inventoryControllerImpl struct {
	inventoryService service.InventoryService
}

func NewInventoryController(inventoryService service.InventoryService) InventoryController {
	return &inventoryControllerImpl{inventoryService: inventoryService}
}

func (c *inventoryControllerImpl) Route(app *fiber.App) {
	group := app.Group("/api/inventories")
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
}

func (c *inventoryControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.inventoryService.FindById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *inventoryControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.inventoryService.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Errors: err.Error(),
		})
	}

	if responses == nil {
		responses = []web.InventoryResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}