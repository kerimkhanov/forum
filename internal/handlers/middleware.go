package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/internal/models"
)

// type appAhandler string

// var beka appAhandler = "2134"

// func (h *Handler) Middleware(next http.Handler) http.Handler {
// 	/*

// 	 */
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		c, err := r.Cookie("session_token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				// http.Redirect(w, r, "/auth", http.StatusSeeOther)
// 				h.ErrorPageHandle(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 				return
// 			}
// 			return
// 		}
// 		user, err := h.service.GetUserWithSession(c.Value)
// 		if err != nil {
// 			fmt.Println("Here2")

// 			h.ErrorPageHandle(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			return
// 		}
// 		if user.TimeSessions.Before(time.Now()) {
// 			if err := h.service.DeleteUserSession(c.Value); err != nil {
// 				fmt.Println("Here3")

// 				h.ErrorPageHandle(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 				return
// 			}
// 			fmt.Println("Here4")
// 			h.ErrorPageHandle(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			return
// 		}
// 		userKey := "user"
// 		ctx := context.WithValue(r.Context(), userKey, user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) models.Users {
// 	c, err := r.Cookie("session_token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			return models.Users{}
// 		}
// 		h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return models.Users{}
// 	}
// 	user, err := h.service.GetUserWithSession(c.Value)
// 	if err != nil {
// 		return models.Users{}
// 	}
// 	if user.ExpiresAt.Before(time.Now()) {
// 		if err := h.service.DeleteUserSession(c.Value); err != nil {
// 			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return models.Users{}
// 		}
// 		return models.Users}
// 	}
// 	return user
// }

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) models.Users {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return models.Users{}
		}
		fmt.Println("userIdentity2")
		h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return models.Users{}
	}
	user, err := h.service.GetUserWithSession(c.Value)
	if err != nil {
		return models.Users{}
	}
	fmt.Printf("\nuser.TimeSessions ---> %s \n", user.TimeSessions)
	if user.TimeSessions.Before(time.Now()) {
		if err := h.service.DeleteUserSession(c.Value); err != nil {
			fmt.Println("userIdentity3")
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return models.Users{}
		}
		return models.Users{}
	}
	return user
}

// func (h *Handler) Middleware(next http.Handler) http.Handler {
// 	/*
// 	 */

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("watch me")
// 	})
// }

// func (h *Handler) Printer() {
// 	fmt.Println("Check print")
// }
