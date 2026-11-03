package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type StockAdjustmentController interface {
	Route(app *fiber.App, authMiddleware ...fiber.Handler)
}

type stockAdjustmentControllerImpl struct {
	stockAdjustmentService service.StockAdjustmentService
}

func NewStockAdjustmentController(stockAdjustmentService service.StockAdjustmentService) StockAdjustmentController {
	return &stockAdjustmentControllerImpl{stockAdjustmentService: stockAdjustmentService}
}

func (c *stockAdjustmentControllerImpl) Route(app *fiber.App, authMiddleware ...fiber.Handler) {
	group := app.Group("/api/stock-adjustments")
	for _, mw := range authMiddleware {
		group.Use(mw)
	}
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
}

func (c *stockAdjustmentControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.StockAdjustmentCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	userIDValue := ctx.Locals("user_id")
	var userID string
	if userIDValue != nil {
		userID, _ = userIDValue.(string)
	}
	if userID == "" {
		userID = ctx.Get("X-User-Id")
	}
	if userID == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "user id not found in context or X-User-Id header")
	}

	response, err := c.stockAdjustmentService.Create(ctx.Context(), request, userID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})
}

func (c *stockAdjustmentControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.stockAdjustmentService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *stockAdjustmentControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.stockAdjustmentService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	if responses == nil {
		responses = []web.StockAdjustmentResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
