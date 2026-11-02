package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type WarehouseController interface {
	Route(app *fiber.App)
}

type warehouseControllerImpl struct {
	warehouseService service.WarehouseService
}

func NewWarehouseController(warehouseService service.WarehouseService) WarehouseController {
	return &warehouseControllerImpl{warehouseService: warehouseService}
}

func (c *warehouseControllerImpl) Route(app *fiber.App) {
	group := app.Group("/api/warehouses")
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
	group.Put("/:id", c.update)
	group.Delete("/:id", c.delete)
}

func (c *warehouseControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.WarehouseCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
	}

	response, err := c.warehouseService.Create(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
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
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
	}

	request.ID = ctx.Params("id")

	response, err := c.warehouseService.Update(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Errors: err.Error(),
		})
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
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Errors: err.Error(),
		})
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

func (c *warehouseControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.warehouseService.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Errors: err.Error(),
		})
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