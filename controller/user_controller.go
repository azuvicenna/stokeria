package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Route(app *fiber.App, authMiddleware ...fiber.Handler)
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

func (c *userControllerImpl) Route(app *fiber.App, authMiddleware ...fiber.Handler) {
	// 1. Endpoint publik untuk login
	app.Post("/api/login", c.login)

	// 2. Endpoint manajemen users (terproteksi jika authMiddleware diberikan)
	group := app.Group("/api/users")
	for _, mw := range authMiddleware {
		group.Use(mw)
	}
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
	group.Put("/:id", c.update)
	group.Delete("/:id", c.delete)
}

func (c *userControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.UserCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	response, err := c.userService.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})
}

func (c *userControllerImpl) update(ctx *fiber.Ctx) error {
	var request web.UserUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	request.ID = ctx.Params("id")

	response, err := c.userService.Update(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *userControllerImpl) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.userService.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "User deleted successfully",
	})
}

func (c *userControllerImpl) findById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := c.userService.FindById(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *userControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.userService.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	if responses == nil {
		responses = []web.UserResponse{}
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}

func (c *userControllerImpl) login(ctx *fiber.Ctx) error {
	var request web.UserLoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	response, err := c.userService.Login(ctx.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
