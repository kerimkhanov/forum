package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"forum/internal/models"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	user := h.userIdentity(w, r)
	if reflect.DeepEqual(user, models.Users{}) {
		fmt.Println("handler.PostCreate - not auth error: ", r.URL.Path)
		h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Println("GET method")
		if err := h.tmpl.ExecuteTemplate(w, "createPost.html", ""); err != nil {
			fmt.Println("handler.CreatePost - ExecuteTemplate Error")
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	case "POST":
		fmt.Println("POst enter")
		err := r.ParseForm()
		if err != nil {
			fmt.Println("err")
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		title, ok := r.Form["title"]
		if !ok {
			fmt.Println("title")
			h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		body, ok := r.Form["body"]
		if !ok {
			fmt.Println("body")
			h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		tags, ok := r.Form["tags"]
		if !ok {
			fmt.Println("tags")
			h.ErrorPageHandle(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		fmt.Println("CreatePost Here  - - - >")
		tagsArr := strings.Split(tags[0], " ")
		fmt.Println("\n "+title[0], body[0], user.Login, strings.Split(tags[0], " "))
		fmt.Println("CreatePost Here  <--")

		// err = h.service.CreatePost(title[0], body[0], user.Login, tagsArr)
		fmt.Println("pointer ===", h.service)
		err = h.service.CreatePost(title[0], body[0], user.Login, tagsArr)

		if err != nil {
			fmt.Println("handler.CreatePost - h.service.CreatePost")
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.ErrorPageHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) PostById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ne zahodit")
	if r.URL.Path != "/posts/view" {
		fmt.Println("handler.PostById -r.URL.Path error: ", r.URL.Path)
		h.ErrorPageHandle(w, "statusNotFound", http.StatusNotFound)
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("handler.PostById -strconv.Atoi error: ", err)
	}
	var OnePost models.OnePost

	OnePost.Users = h.userIdentity(w, r)
	switch r.Method {
	case "GET":
		OnePost.Posts, err = h.service.GetPostById(id)
		// comments := []models.Comments{{Comment: "Hasan 4mo"}, {Comment: "Andrei krut"}, {Comment: "Alisher Top4eg"}}
		// OnePost.Comments = append(OnePost.Comments, comments...)
		if err != nil {
			fmt.Println("handler.PostById - GetPostById error: ", err)
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		postLikes, err := h.service.GetPostLikes(id)
		if err != nil {
			fmt.Println("handler.PostById - GetPostById error2: ", err)
			h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
			return
		}
		postDislikes, err := h.service.GetPostDislikes(id)
		if err != nil {
			fmt.Println("handler.PostById - GetPostById error3: ", err)
			h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
			return
		}
		commentsLikes, err := h.service.GetCommentLikes(id, OnePost.Login)
		if err != nil {
			fmt.Println("handler.PostById - GetPostById error4: ", err)

			h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
			return
		}
		commentsDislikes, err := h.service.GetCommnetDislikes(id, OnePost.Login)
		if err != nil {
			fmt.Println("handler.PostById - GetPostById error5: ", err)

			h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i := 0; i < len(OnePost.Comments); i++ {
			OnePost.Comments[i].Likes += len(commentsLikes[OnePost.Comments[i].Comment_id])
		}
		for i := 0; i < len(OnePost.Comments); i++ {
			OnePost.Comments[i].Dislikes += len(commentsDislikes[OnePost.Comments[i].Comment_id])
		}
		fmt.Println(commentsLikes, commentsDislikes)
		info := models.Info{
			Posts:            OnePost.Posts,
			PostLikes:        postLikes,
			PostDislikes:     postDislikes,
			Users:            OnePost.Users,
			Comments:         OnePost.Comments,
			CommentsLikes:    commentsLikes,
			CommentsDislikes: commentsDislikes,
		}
		if err := h.tmpl.ExecuteTemplate(w, "postPages.html", info); err != nil {
			fmt.Println("Alish pidr")
			fmt.Println(err, "------------------")
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Println(3)
	case "POST":
		if err := r.ParseForm(); err != nil {
			h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		comment, ok := r.Form["Comment"]
		fmt.Printf("here comment: --> %s", comment)
		if !ok {
			h.ErrorPageHandle(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		err := h.service.AddComment(id, comment[0], OnePost.Users.Login)
		if err != nil {
			h.ErrorPageHandle(w, http.StatusText(http.StatusInternalServerError), http.StatusSeeOther)
		}
		http.Redirect(w, r, fmt.Sprintf("/posts/view?id=%d", id), http.StatusSeeOther)
		return
	default:
		h.ErrorPageHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
