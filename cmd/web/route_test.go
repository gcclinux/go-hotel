package main

import (
	"fmt"
	"myapp/internal/config"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
