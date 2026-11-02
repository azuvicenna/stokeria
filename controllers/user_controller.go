package controller

import (
	"stokeria/model/web"
	"stokeria/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Route(app *fiber.App)
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

func (c *userControllerImpl) Route(app *fiber.App) {
	group := app.Group("/api/users")
	group.Post("/", c.create)
	group.Get("/", c.findAll)
	group.Get("/:id", c.findById)
	group.Put("/:id", c.update)
	group.Delete("/:id", c.delete)

	// Route tambahan untuk login. Jika Anda akan menggunakan JWT,
	// endpoint ini sebaiknya diletakkan di controller tersendiri (misal AuthController).
	app.Post("/api/login", c.login)
}

func (c *userControllerImpl) create(ctx *fiber.Ctx) error {
	var request web.UserCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
	}

	response, err := c.userService.Create(ctx.Context(), request)
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

func (c *userControllerImpl) update(ctx *fiber.Ctx) error {
	var request web.UserUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
	}

	request.ID = ctx.Params("id")

	response, err := c.userService.Update(ctx.Context(), request)
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

func (c *userControllerImpl) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.userService.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Errors: err.Error(),
		})
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

func (c *userControllerImpl) findAll(ctx *fiber.Ctx) error {
	responses, err := c.userService.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Errors: err.Error(),
		})
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
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: err.Error(),
		})
	}

	response, err := c.userService.Login(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Errors: err.Error(),
		})
	}

	// Di aplikasi nyata, Anda mungkin ingin membuat JWT Token di sini
	// dan mengembalikannya alih-alih hanya mereturn data User.
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}