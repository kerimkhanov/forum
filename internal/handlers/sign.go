package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"forum/internal/models"
	"forum/internal/service"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/logout" {
		fmt.Println("handler.logout - r.URL.Path error: ", r.URL.Path)
		h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = h.service.DeleteUserSession(c.Value)
	if err != nil {
		fmt.Println("logout - DeleteUserSession: ", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.URL.Path != "/auth" {
		h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user := h.userIdentity(w, r)
	if !reflect.DeepEqual(user, models.Users{}) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	switch r.Method {
	case "GET":
		fmt.Println("GET method")
		if err := h.tmpl.ExecuteTemplate(w, "auth.html", ""); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	case "POST":
		email, ok := r.Form["email"]
		if !ok {
			fmt.Println("wtf@1?")
			h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		password, ok := r.Form["password"]
		if !ok {
			fmt.Println("wtf@2?")
			h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		fmt.Println("POst enter")
		user, err := h.service.CreateSession(email[0], password[0])
		fmt.Printf("\n user seesion %s \n", user.Session_token)
		if err != nil {
			fmt.Printf("handler - handlers.go - MainPage %v", err)
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   user.Session_token,
			Expires: time.Now().Add(120 * time.Hour),
		})
		http.Redirect(w, r, "/", 301)
	default:
		h.ErrorPageHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/SignUp" {
		fmt.Println("handler.SignUp - r.URL.Path error: ", r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user := h.userIdentity(w, r)
	fmt.Printf("\nUserIdentity result --> %s\n", user)
	if !reflect.DeepEqual(user, models.Users{}) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	switch r.Method {
	case "GET":
		fmt.Println("enter GET")
		if err := h.tmpl.ExecuteTemplate(w, "SignUp.html", nil); err != nil {
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case "POST":
		fmt.Println("enter POST")
		pass := service.PasswordHash(r.FormValue("password"))
		user = models.Users{
			Login:    r.FormValue("Login"),
			Email:    r.FormValue("email"),
			Password: pass,
		}
		err := h.service.AddUsers(user)
		if err != nil {
			fmt.Printf("handlers - handlers.go - signup - AddUser: - %v", err)
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Printf("\n %s Login = %s Email = %s Password = %s\n", "Handler After h.serbice.AddUsers", user.Login, user.Email, user.Password)
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
	default:
		h.ErrorPageHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
