package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/forms"
	"myapp/internal/models"
	"myapp/internal/render"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the home page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// send data to teplate
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

// General is the General's Quarters page handler
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Major is the Major's Suite page handler
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page used by the handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Availability is the Availability search page used by the handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the PostAvailability search page used by the handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// And that directive says, take whatever you see inside the parentheses, hear this string and convert
	// it to a slice of bytes.
	// w.Write([]byte("Post to search availability"))

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send json
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Println(err)
	}

	//log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// MkReservation renders the make a reservation page and display form
func (m *Repository) MkReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostMkReservation handles the posting of a reservation form
func (m *Repository) PostMkReservation(w http.ResponseWriter, r *http.Request) {

}
