package app

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"stokeria/exception"

	"github.com/gofiber/fiber/v2"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	api := app.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Stokeria service is running",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if result["status"] != "success" {
		t.Errorf("Expected status 'success', got %v", result["status"])
	}
}

func TestNotFoundHandler(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	req := httptest.NewRequest(http.MethodGet, "/api/nonexistent", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var errResp exception.ErrorResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		t.Fatalf("Failed to parse error response: %v", err)
	}

	if errResp.Code != http.StatusNotFound {
		t.Errorf("Expected error code 404, got %d", errResp.Code)
	}
}
