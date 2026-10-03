package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doHello(t *testing.T, target string) Greeting {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	handleHello(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}

	var g Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &g); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	return g
}

func TestHelloWithName(t *testing.T) {
	g := doHello(t, "/hello?name=Alice")
	if g.Name != "Alice" {
		t.Errorf("name = %q, want %q", g.Name, "Alice")
	}
	if g.Message != "Hello, Alice!" {
		t.Errorf("message = %q, want %q", g.Message, "Hello, Alice!")
	}
}

func TestHelloWithoutName(t *testing.T) {
	g := doHello(t, "/hello")
	if g.Name != "" {
		t.Errorf("name = %q, want empty", g.Name)
	}
	if g.Message != "Hello, World!" {
		t.Errorf("message = %q, want %q", g.Message, "Hello, World!")
	}
}

func TestHelloWithEmptyName(t *testing.T) {
	g := doHello(t, "/hello?name=")
	if g.Name != "" {
		t.Errorf("name = %q, want empty", g.Name)
	}
	if g.Message != "Hello, World!" {
		t.Errorf("message = %q, want %q", g.Message, "Hello, World!")
	}
}
