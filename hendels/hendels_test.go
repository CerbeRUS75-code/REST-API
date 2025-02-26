package hendels

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetHendler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/messages", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := GetHendler(c); err != nil {
		t.Errorf("Error was not expected: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rec.Code)
	}
}

func TestPostHendler(t *testing.T) {
	e := echo.New()
	message := map[string]string{"content": "Hello, World!"}
	jsonMessage, _ := json.Marshal(message)
	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewBuffer(jsonMessage))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := PostHendler(c); err != nil {
		t.Errorf("Error was not expected: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", rec.Code)
	}
}

func TestDeleteHendler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/messages/1", nil) // Предполагаем, что ID = 1
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := DeleteHendler(c); err != nil {
		t.Errorf("Error was not expected: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rec.Code)
	}
}

func TestPatchHendler(t *testing.T) {
	e := echo.New()
	message := map[string]string{"content": "Updated message"}
	jsonMessage, _ := json.Marshal(message)
	req := httptest.NewRequest(http.MethodPatch, "/messages/1", bytes.NewBuffer(jsonMessage)) // Предполагаем, что ID = 1
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := PatchHendler(c); err != nil {
		t.Errorf("Error was not expected: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rec.Code)
	}
}

// Негативные тесты
func TestPostHendlerInvalidJSON(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewBuffer([]byte(`{invalid json}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := PostHendler(c); err == nil {
		t.Error("Expected an error, got none")
	}
}

func TestDeleteHendlerNonExistentID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/messages/999", nil) // Предполагаем, что ID = 999 не существует
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := DeleteHendler(c); err != nil {
		t.Errorf("Error was not expected: %v", err)
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", rec.Code)
	}
}

// Добавьте аналогичные тесты для PatchHendler и DeleteHendler
