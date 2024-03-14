package json_helpers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func TestReadJSON(t *testing.T) {
	// Create a new http.Request with a JSON body
	req := httptest.NewRequest("GET", "http://example.com/foo", bytes.NewBuffer([]byte(`{"field1":"test", "field2":123}`)))
	w := httptest.NewRecorder()

	// Create a new TestStruct to hold the decoded JSON
	data := &TestStruct{}

	// Call the ReadJSON function
	err := ReadJSON(w, req, data)

	// Check for errors
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check that the decoded JSON matches the expected values
	if data.Field1 != "test" || data.Field2 != 123 {
		t.Errorf("expected field1=test and field2=123, got field1=%s and field2=%d", data.Field1, data.Field2)
	}
}
func TestWriteJSON(t *testing.T) {
	// Create a new http.ResponseWriter
	w := httptest.NewRecorder()

	// Create a new TestStruct with some data
	data := &TestStruct{
		Field1: "test",
		Field2: 123,
	}
	fmt.Println(data)

	// Call the WriteJSON function
	err := WriteJSON(w, http.StatusOK, data)

	// Check for errors
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check that the response body matches the expected JSON
	expected := `{"field1":"test","field2":123}`
	if w.Body.String() != expected {
		t.Errorf("expected body to be %s, got %s", expected, w.Body.String())
	}
}

func TestErrorJSON(t *testing.T) {
	// Create a new http.ResponseWriter
	w := httptest.NewRecorder()

	// Call the ErrorJSON function
	err := fmt.Errorf("page not found")
	err = ErrorJSON(w, err, http.StatusNotFound)

	// Check for errors
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	fmt.Println(w.Body)
	// Check that the response body matches the expected JSON
	expected := `{"error":true,"message":"page not found"}`
	if w.Body.String() != expected {
		t.Errorf("expected body to be %s, got %s", expected, w.Body.String())
	}
}
