package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"forum/internal/models"
	"forum/internal/service"

	uuid "github.com/satori/go.uuid"
)

type Handler struct {
	service *service.Service
	tmpl    *template.Template
}

// type Handler interface {
// 	Register(router *http.ServeMux)
// }

func NewHandler(service *service.Service) *Handler {
	tmp, _ := template.ParseGlob("./templates/*.html")
	return &Handler{
		service: service,
		tmpl:    tmp,
	}
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.MainPage)
	router.HandleFunc("/auth", h.Authorization)
	router.HandleFunc("/SignUp", h.SignUp)
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err := h.tmpl.ExecuteTemplate(w, "index.html", ""); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

//  При авторизации пользователя ты сравниваешь пароль с базой данных в которй он захэширован. Расшехируешь пароль в базе и справниваешь его со своим
//  Если данные правильные создаешь сессию для этого пользователя и даешь ему куки
func (h *Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" {
		h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// user := h.userIdentity(w, r)
	// if !reflect.DeepEqual(user, models.Users{}) {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// }
	switch r.Method {
	case "GET":
		fmt.Println("GET method")
		if err := h.tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	case "POST":
		fmt.Println("POst enter")
		// user, err := h.service.UserByEmail(r.FormValue("email"))
		// if err != nil {
		// 	fmt.Errorf("Handlers -> Authorization -> POST", err.Error())
		// 	return
		// }
		// err = service.CorrectAuth(user, r.FormValue("email"), r.FormValue("password"))
		// if err != nil {
		// 	fmt.Errorf("Handlers -> Authorization -> CorrectAuth", err)
		// 	if err := h.tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
		// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		// 	}
		// }
		// if err != nil {
		// 	if err := h.tmpl.ExecuteTemplate(w, "index.html", ""); err != nil {
		// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		// 	}
		// }
		// set cookies
		uuid := uuid.NewV4().String()
		// save uuid in db connect with user
		user_id, err := strconv.Atoi(user.Id)
		if err != nil {
			log.Println("error creating user", err)
			if err := h.tmpl.ExecuteTemplate(w, "index.html", ""); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
		user, err := h.service.CreateSession(user_id)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   user.Session_token,
			Expires: time.Now().Add(120 * time.Hour),
		})
		w.WriteHeader(200)
		return
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/SignUp" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// user := h.userIdentity(w, r)
	// if !reflect.DeepEqual(user, models.Users{}) {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// }
	// storage.AddUsers(&sql.DB{}, r.FormValue("login"), r.FormValue("email"), r.FormValue("password"))
	pass := service.PasswordHash(r.FormValue("password"))
	user := models.Users{
		Login:    r.FormValue("Login"),
		Email:    r.FormValue("email"),
		Password: pass,
	}
	err := h.service.AddUsers(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n %s Login = %s Email = %s Password = %s\n", "Handler After h.serbice.AddUsers", user.Login, user.Email, user.Password)
	if err := h.tmpl.ExecuteTemplate(w, "SignUp.html", ""); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handler) ErrorPageHandle(w http.ResponseWriter, response string, code int) {
	err := models.Err{
		Code:     code,
		Response: response,
	}
	w.WriteHeader(code)
	if err := h.tmpl.ExecuteTemplate(w, "error.html", err); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
