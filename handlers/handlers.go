package handlers

import (
	"fmt"
	"forum/internal/appMiddleware"
	"forum/internal/models"
	"forum/internal/repository"
	internal "forum/internal/repository"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
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
	router.HandleFunc("/", appMiddleware.Middleware(h.MainPage))
	router.HandleFunc("/auth", appMiddleware.Middleware(h.Authorization))
	router.HandleFunc("/Regis", appMiddleware.Middleware(h.Regis))
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
		fmt.Println("GET method")
		if err := tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	case "POST":
		fmt.Println("POst enter")
		user, err := h.service.UserByEmail(r.FormValue("email"))
		if err != nil {
			fmt.Errorf("Handlers -> Authorization -> POST", err.Error())
			return
		}
		err = repository.CorrectAuth(user, r.FormValue("email"), r.FormValue("password"))
		if err != nil {
			fmt.Errorf("Handlers -> Authorization -> CorrectAuth", err)
			if err := tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
		if err != nil {
			if err := tmpl.ExecuteTemplate(w, "index.html", ""); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
		// set cookies

		uuid := uuid.NewV4().String()
		// save uuid in db connect with user
		user_id, err := strconv.Atoi(user.Id)
		if err != nil {
			log.Println("error creating user", err)
			if err := tmpl.ExecuteTemplate(w, "index.html", ""); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
		h.service.CreateSession(user_id, uuid, time.Now())
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   uuid,
			Expires: time.Now().Add(120 * time.Hour),
		})
		w.WriteHeader(200)
		return
		// w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	}
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
	pass := repository.PasswordHash(r.FormValue("password"))
	user := models.Users{
		Login:    r.FormValue("Login"),
		Email:    r.FormValue("email"),
		Password: pass,
	}
	h.service.AddUsers(user)

	// fmt.Print(User.Password)
	if err := tmpl.ExecuteTemplate(w, "Regis.html", ""); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
