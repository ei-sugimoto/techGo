package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ei-sugimoto/techGO/pkg/middleware"
)

func TestRecovery(t *testing.T) {
	testW := httptest.NewRecorder()

	testR, err := http.NewRequest("GET", "/panic", nil)
	if err != nil {
		t.Fatal(err)
	}

	testH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("panic test")
	})

	middleware.Recovery(testH).ServeHTTP(testW, testR)

	if status := testW.Result().StatusCode; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}
