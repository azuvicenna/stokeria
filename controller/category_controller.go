package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Route(app *fiber.App, authMiddleware ...fiber.Handler)
}

type categoryControllerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryControllerImpl{categoryService: categoryService}
}

func (c *categoryControllerImpl) Route(app *fiber.App, authMiddleware ...fiber.Handler) {
	group := app.Group("/api/categories")
	for _, mw := range authMiddleware {
		group.Use(mw)
	}
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
	group.Put("/:id", c.update)
	group.Delete("/:id", c.delete)
}

func (c *categoryControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.CategoryCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	response, err := c.categoryService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})
}

func (c *categoryControllerImpl) update(ctx *fiber.Ctx) error {
	var request web.CategoryUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	request.ID = ctx.Params("id")

	response, err := c.categoryService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *categoryControllerImpl) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.categoryService.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Category deleted successfully",
	})
}

func (c *categoryControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.categoryService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *categoryControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.categoryService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	if responses == nil {
		responses = []web.CategoryResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
