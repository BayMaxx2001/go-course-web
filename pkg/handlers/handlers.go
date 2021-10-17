package handlers

import (
	"net/http"

	"github.com/BayMaxx2001/go-course-web/pkg/config"
	"github.com/BayMaxx2001/go-course-web/pkg/models"
	"github.com/BayMaxx2001/go-course-web/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository iss the Repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// home is the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
