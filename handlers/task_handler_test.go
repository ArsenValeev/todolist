package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "study/handlers"
)

func TestCreateTaskHandler(t *testing.T) {
    // Создаем тестовый запрос
    body := `{"title": "Test", "description": ""}`
    req := httptest.NewRequest("POST", "/tasks", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-ID", "1")
    
    // Записываем ответ
    rr := httptest.NewRecorder()
    handler := handlers.NewtaskHandlerMySQL()
    
    // Вызываем handler
    handler.CreateTaskDB(rr, req)
    
    // Проверяем статус код
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status 200, got %d", status)
    }
    
    // Проверяем тело ответа
    expected := `"title":"Test"`
    if !strings.Contains(rr.Body.String(), expected) {
        t.Errorf("Response should contain %s", expected)
    }
}