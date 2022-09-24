package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) commentDislike(w http.ResponseWriter, r *http.Request) {
	user := h.userIdentity(w, r)
	comment_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("handler.CommentDislike -strconv.Atoi error : %v", err)
	}
	post_id, err := h.service.GetPostIdByComment(comment_id, user.Login)
	if err != nil {
		fmt.Println("handler.like.CommentDislike - commentDislike: %v", err)
	}
	if err := h.service.AddCommentDislike(comment_id, user.Login); err != nil {
		fmt.Println("error here")
		h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/view?id=%d", post_id), http.StatusSeeOther)
}

func (h *Handler) commentLike(w http.ResponseWriter, r *http.Request) {
	user := h.userIdentity(w, r)
	comment_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("handler.CommentLike - strconv.Atoi error: %v", err)
		return
	}
	post_id, err := h.service.GetPostIdByComment(comment_id, user.Login)
	if err != nil {
		fmt.Println("handler.like.commentLike - commentLike: %v", err)
		return
	}

	err = h.service.AddCommentLike(comment_id, user.Login)
	if err != nil {
		fmt.Println("handler.like.commentLike - addCommentLike: %v", err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/view?id=%d", post_id), http.StatusSeeOther)
}

func (h *Handler) PostDislike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here some2")
	user := h.userIdentity(w, r)
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("Handler.PostLike - strconv.Atoi error: %v", err)
		return
	}
	if err := h.service.DislikePost(post_id, user.Login); err != nil {
		// if errors.Is(err, sql.ErrNoRows) {
		// 	h.ErrorPageHandle(w, err.Error(), http.StatusNotFound)
		// 	return
		// }
		h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/posts/view?id=%d", post_id), http.StatusSeeOther)
}

func (h *Handler) PostLike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here some")
	user := h.userIdentity(w, r)
	post_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Println(post_id)
	if err != nil {
		fmt.Println("Handler.PostLike - strconv.Atoi error: %v", err)
		return
	}
	// if r.Method != http.MethodPost {
	// 	h.ErrorPageHandle(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }
	if err = h.service.LikePost(post_id, user.Login); err != nil {
		fmt.Println("sam petuh")
		h.ErrorPageHandle(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/view?id=%d", post_id), http.StatusSeeOther)
}
