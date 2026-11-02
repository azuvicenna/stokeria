package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type OutboundTransactionController interface {
	Route(app *fiber.App)
}

type outboundTransactionControllerImpl struct {
	outboundTransactionService service.OutboundTransactionService
}

func NewOutboundTransactionController(outboundTransactionService service.OutboundTransactionService) OutboundTransactionController {
	return &outboundTransactionControllerImpl{outboundTransactionService: outboundTransactionService}
}

func (c *outboundTransactionControllerImpl) Route(app *fiber.App) {
	group := app.Group("/api/outbound-transactions")
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
}

func (c *outboundTransactionControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.OutboundTransactionCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
	}

	userIDValue := ctx.Locals("user_id")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Errors: "user id not found in context",
		})
	}

	userID, ok := userIDValue.(string)
	if !ok || userID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Errors: "invalid user id format",
		})
	}

	response, err := c.outboundTransactionService.Create(ctx.Context(), request, userID)
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

func (c *outboundTransactionControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.outboundTransactionService.FindById(ctx.Context(), id)
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

func (c *outboundTransactionControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.outboundTransactionService.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Errors: err.Error(),
		})
	}

	if responses == nil {
		responses = []web.OutboundTransactionResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}