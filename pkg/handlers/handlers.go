package handlers

import (
	"github.com/tanc7/hotelbookings/models"
	"github.com/tanc7/hotelbookings/pkg/config"
	"github.com/tanc7/hotelbookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new respoitory
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// The first arg specifies a "receiver" which is the instance of Repository struct, you need to set it to render.repo.repository in main because it specifies a receiver now.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

//	fmt.Fprintf(w, "This is the homepage")
//}

//About is /about path
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	// send data to template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

//sum := addValues(2, 2)
//fmt.Fprintf(w, "This is the About page")
////fmt.Fprintf(w, "The sum of these values is %d", sum)
//_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 = %d", sum))

//func AddValues(x, y int) (int, error) {
//	//var sum int
//	//sum = x + y
//	var sum int
//	sum = x + y
//	return sum, nil
//}
//addValues returns x + y
//func addValues(x, y int) int {
//	return x + y
//}
//
//func Divide(w http.ResponseWriter, r *http.Request) {
//	f, err := divideValues(100.0, 0.0)
//	if err != nil {
//		fmt.Fprintf(w, "Cannot by divide by 0")
//		return
//	}
//
//	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
//}
//
//func divideValues(x, y float32) (float32, error) {
//	if y <= 0 {
//		err := errors.New("Cannot divide by zero")
//		return 0, err
//	}
//	result := x / y
//	return result, nil
//}
