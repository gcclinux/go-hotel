package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repositorty
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("handlers.go: app.InProduction", m.App.InProduction)

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the home page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// send data to teplate
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

// General is the General's Quarters page handler
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Major is the Major's Suite page handler
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page used by the handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Availability is the Availability search page used by the handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// MkReservation is the book now page used by the handler
func (m *Repository) MkReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "mk-reservation.page.tmpl", &models.TemplateData{})
}
