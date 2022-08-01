package handlers

import (
	"github.com/xor0x/bookings/pkg/config"
	"net/http"
	"github.com/xor0x/bookings/pkg/models"
	"github.com/xor0x/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// newRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is a function that handles the / route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is a function that handles the /about route
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
stringMap := make(map[string]string)
stringMap["test"] = "Hello World"

remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")


stringMap["remote_ip"] = remoteIP

render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
