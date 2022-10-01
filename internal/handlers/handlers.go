package handlers

import (
	"errors"
	"fmt"
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
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/", h.MainPage)
	router.HandleFunc("/auth", h.Authorization)
	router.HandleFunc("/SignUp", h.SignUp)
	router.HandleFunc("/posts/create", h.CreatePost)
	router.HandleFunc("/posts/view", h.PostById)
	router.HandleFunc("/auth/logout", h.Logout)
	router.HandleFunc("/posts/like", h.PostLike)
	router.HandleFunc("/posts/dislike", h.PostDislike)
	router.HandleFunc("/comments/like", h.commentLike)
	router.HandleFunc("/comments/dislike", h.commentDislike)
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		var Allpost models.Allpost
		var err error
		Allpost.Users = h.userIdentity(w, r)
		if len(r.URL.Query()) == 0 {
			Allpost.Posts, err = h.service.GetAllPosts()
			if err != nil {
				fmt.Println(err)
				h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		} else {
			Allpost.Posts, err = h.service.GetAllPostsBy(Allpost.Users, r.URL.Query())
			if err != nil {
				if errors.Is(err, errors.New("user does not exist or password incorrect")) {
					h.ErrorPageHandle(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
			}
		}
		user := struct {
			Users models.Users
			Posts []models.Posts
		}{
			Allpost.Users,
			Allpost.Posts,
		}
		if err := h.tmpl.ExecuteTemplate(w, "index.html", user); err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		h.ErrorPageHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
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
