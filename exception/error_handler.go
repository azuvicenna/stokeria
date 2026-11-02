package exception

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct{
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Errors map[string]string `json:"errors"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	status := "INTERNAL_SERVER_ERROR"
	errorMessages := make(map[string]string)

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		status = strings.ToUpper(strings.ReplaceAll(http.StatusText(code), " ", "_"))
		errorMessages["server"] = fiberErr.Message
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		code = fiber.StatusBadRequest
		status = "BAD_REQUEST"
		for _, fieldError := range validationErrs {
			// Menggunakan Field() sebagai key dan ActualTag() sebagai deskripsi error
			errorMessages[fieldError.Field()] = "Validation failed on tag: " + fieldError.ActualTag()
		}
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	errorMessages["general"] = err.Error()
	return writeJSONResponse(ctx, code, status, errorMessages)
}

func writeJSONResponse(ctx *fiber.Ctx, code int, status string, errors map[string]string) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return ctx.Status(code).JSON(ErrorResponse{
		Code:   code,
		Status: status,
		Errors: errors,
	})
}