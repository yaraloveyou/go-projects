package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_formHandler(t *testing.T) {
	formData := strings.NewReader("name=John&email=john@example.com&tel=123456789")
	req, err := http.NewRequest("POST", "/form", formData)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(formHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Form name=John email=john@example.com tel=123456789"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func Test_formHandlerParseForm(t *testing.T) {
	req, err := http.NewRequest("POST", "/form", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(formHandler)

	handler.ServeHTTP(rr, req)

	// Проверяем, что возвращается ошибка парсинга формы
	if !strings.Contains(rr.Body.String(), "ParseForm() err") {
		t.Errorf("handler should have returned a form parse error")
	}
}

func Test_helloHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/wrongpath", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func Test_helloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	excepted := "Hello World!"

	if rr.Body.String() != excepted {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), excepted)
	}
}
