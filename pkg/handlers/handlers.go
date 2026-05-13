package handlers

import (
	"net/http"

	"github.com/rohanv8676/bookings/pkg/config"
	"github.com/rohanv8676/bookings/pkg/models"
	"github.com/rohanv8676/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About Page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some business logic here, which gives me some data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again mofo."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// Send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
