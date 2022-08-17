package handlers

import (
	"html/template"
	"net/http"

	"forum/internal/models"
	"forum/internal/service"
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
