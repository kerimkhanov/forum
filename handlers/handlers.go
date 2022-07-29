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
		fmt.Println("GET method")
		if err := tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	case "POST":
		fmt.Println("POst enter")
		
		user, err := h.service.UserByEmail(email)
		if err != nil {
			fmt.Errorf("Handlers -> Authorization -> POST", err.Error())
			return
		}
		fmt.Println(err)
		err = CorrectAuth(user2, r.FormValue("Email"), r.FormValue("Password"))
		fmt.Println(err)
		if err == nil {
			if err := tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		// if err == nil {
		// http.Redirect(w, r, "/", http.StatusFound)
		// }

	}
}

func CorrectAuth(user models.Users, Email, Password string) error {
	fmt.Println(1)
	count := 0
	fmt.Println(user)
	fmt.Println(Password)
	if Email == user.Email {
		count++
	}
	if Password == user.Password {
		count++
	}
	if count != 2 {
		fmt.Println(2)
		return errors.New("Invalid username/password")
	}
	fmt.Println(3)

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
