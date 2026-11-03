package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type WarehouseController interface {
	Route(app *fiber.App, authMiddleware ...fiber.Handler)
}

type warehouseControllerImpl struct {
	warehouseService service.WarehouseService
}

func NewWarehouseController(warehouseService service.WarehouseService) WarehouseController {
	return &warehouseControllerImpl{warehouseService: warehouseService}
}

func (c *warehouseControllerImpl) Route(app *fiber.App, authMiddleware ...fiber.Handler) {
	group := app.Group("/api/warehouses")
	for _, mw := range authMiddleware {
		group.Use(mw)
	}
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
	group.Put("/:id", c.update)
	group.Delete("/:id", c.delete)
}

func (c *warehouseControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.WarehouseCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	response, err := c.warehouseService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})
}

func (c *warehouseControllerImpl) update(ctx *fiber.Ctx) error {
	var request web.WarehouseUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	request.ID = ctx.Params("id")

	response, err := c.warehouseService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *warehouseControllerImpl) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.warehouseService.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Warehouse deleted successfully",
	})
}

func (c *warehouseControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.warehouseService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *warehouseControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.warehouseService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	if responses == nil {
		responses = []web.WarehouseResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
