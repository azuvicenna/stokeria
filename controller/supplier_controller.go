package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type SupplierController interface {
	Route(app *fiber.App, authMiddleware ...fiber.Handler)
}

type supplierControllerImpl struct {
	supplierService service.SupplierService
}

func NewSupplierController(supplierService service.SupplierService) SupplierController {
	return &supplierControllerImpl{supplierService: supplierService}
}

func (c *supplierControllerImpl) Route(app *fiber.App, authMiddleware ...fiber.Handler) {
	group := app.Group("/api/suppliers")
	for _, mw := range authMiddleware {
		group.Use(mw)
	}
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
	group.Put("/:id", c.update)
	group.Delete("/:id", c.delete)
}

func (c *supplierControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.SupplierCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	response, err := c.supplierService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})
}

func (c *supplierControllerImpl) update(ctx *fiber.Ctx) error {
	var request web.SupplierUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	request.ID = ctx.Params("id")

	response, err := c.supplierService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *supplierControllerImpl) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.supplierService.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Supplier deleted successfully",
	})
}

func (c *supplierControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.supplierService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *supplierControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.supplierService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	if responses == nil {
		responses = []web.SupplierResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
