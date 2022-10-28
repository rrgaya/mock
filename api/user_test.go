package api

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetRandomUser(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	GetRandomUser(w, r)

	got := w.Result()

	if !reflect.DeepEqual(http.StatusOK, got.StatusCode) {
		t.Errorf("Wanted: %d, Got: %d", http.StatusOK, got.StatusCode)
	}

}
