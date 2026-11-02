package exception

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Errors map[string]string `json:"errors"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	status := "INTERNAL_SERVER_ERROR"
	errorMessages := make(map[string]string)

	// 1. Fiber Error (dibuat manual via fiber.NewError)
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		status = strings.ToUpper(strings.ReplaceAll(http.StatusText(code), " ", "_"))
		errorMessages["message"] = fiberErr.Message
		logIfServerError(ctx, code, err)
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	// 2. Go Playground Validator Errors
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		code = fiber.StatusBadRequest
		status = "BAD_REQUEST"
		for _, fieldError := range validationErrs {
			errorMessages[fieldError.Field()] = "Validation failed on tag: " + fieldError.ActualTag()
		}
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	// 3. Database Record Not Found
	if errors.Is(err, gorm.ErrRecordNotFound) || strings.Contains(strings.ToLower(err.Error()), "not found") {
		code = fiber.StatusNotFound
		status = "NOT_FOUND"
		errorMessages["general"] = err.Error()
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	// 4. Duplicate / Insufficient / Invalid Business Rules
	errLower := strings.ToLower(err.Error())
	if strings.Contains(errLower, "already exists") ||
		strings.Contains(errLower, "insufficient stock") ||
		strings.Contains(errLower, "invalid") ||
		strings.Contains(errLower, "duplicate") {
		code = fiber.StatusBadRequest
		status = "BAD_REQUEST"
		errorMessages["general"] = err.Error()
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	// 5. Unauthorized
	if strings.Contains(errLower, "unauthorized") {
		code = fiber.StatusUnauthorized
		status = "UNAUTHORIZED"
		errorMessages["general"] = err.Error()
		return writeJSONResponse(ctx, code, status, errorMessages)
	}

	// 6. Default Internal Server Error
	logIfServerError(ctx, code, err)
	errorMessages["general"] = err.Error()
	return writeJSONResponse(ctx, code, status, errorMessages)
}

func logIfServerError(ctx *fiber.Ctx, code int, err error) {
	if code >= fiber.StatusInternalServerError {
		logrus.WithFields(logrus.Fields{
			"method": ctx.Method(),
			"path":   ctx.Path(),
			"status": code,
		}).Errorf("[SERVER ERROR] %v", err)
	}
}

func writeJSONResponse(ctx *fiber.Ctx, code int, status string, errors map[string]string) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return ctx.Status(code).JSON(ErrorResponse{
		Code:   code,
		Status: status,
		Errors: errors,
	})
}
