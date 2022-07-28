package handlers

import (
	"errors"
	"fmt"
	"forum/internal"
	"forum/internal/models"
	"html/template"
	"net/http"
)

type Handler struct {
	service *internal.Service
}

// type Handler interface {
// 	Register(router *http.ServeMux)
// }

func NewHandler(service *internal.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.MainPage)
	router.HandleFunc("/auth", h.Authorization)
	router.HandleFunc("/Regis", h.Regis)
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", ""); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	// h.serviceService.Register(tmpl)
}

//  При авторизации пользователя ты сравниваешь пароль с базой данных в которй он захэширован. Расшехируешь пароль в базе и справниваешь его со своим
//  Если данные правильные создаешь сессию для этого пользователя и даешь ему куки
func (h *Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/auth.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		if err := tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	case "POST":
		user2 := models.Users{}
		err = h.service.GetUsers(user2)
		if err != nil {
			fmt.Errorf("Handlers -> Authorization -> POST", err.Error())
			return
		}
		err = CorrectAuth(user2, r.FormValue("Login"), r.FormValue("Email"), r.FormValue("Password"))
		if err != nil {
			fmt.Errorf("Handlers -> Authorization -> CorrectAuth", err.Error())
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func CorrectAuth(user models.Users, Login, Email, Password string) error {
	count := 0
	if Login == user.Login {
		count++
	}
	if Email == user.Email {
		count++
	}
	if Password == user.Password {
		count++
	}
	if count != 3 {
		return errors.New("Invalid username/password")
	}
	return nil
}

func (h *Handler) Regis(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Regis" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/Regis.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}
	// storage.AddUsers(&sql.DB{}, r.FormValue("login"), r.FormValue("email"), r.FormValue("password"))
	user := models.Users{
		Login:    r.FormValue("Login"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	h.service.AddUsers(user)

	// fmt.Print(User.Password)
	if err := tmpl.ExecuteTemplate(w, "Regis.html", ""); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
