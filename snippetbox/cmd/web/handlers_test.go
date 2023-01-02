package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pqppq/lets-go/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	rs := rr.Result()

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	body, err := io.ReadAll(rs.Body)
	defer rs.Body.Close()

	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}