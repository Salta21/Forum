package delivery

import (
	"net/http"
	"strconv"

	"forum/web/models"
)

func (h *Handler) home(writer http.ResponseWriter, request *http.Request) { // home page
	if request.URL.Path != "/" && request.URL.Path != "/home" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if userId != 0 {
		http.Redirect(writer, request, "/profile", http.StatusSeeOther)
	}
	switch request.Method {
	case http.MethodGet:
		if err := h.render(writer, "home.html", nil); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := request.ParseForm(); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		b, ok := request.Form["content"]
		if !ok {
			h.errorPage(writer, request, http.StatusBadRequest)
			return
		}
		button := b[0]
		displayallpost, err := h.services.GetAllPostsByCategory(button)
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		u := models.User{
			Post: displayallpost,
		}
		if err = h.render(writer, "home.html", u); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) displayForUnauthorized(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/displayforunauthorized/" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if userId != 0 {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	if request.Method != http.MethodGet {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	post_id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	post, err := h.services.GetPostByPostId(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	countLiked, err := h.services.CountLikedPosts(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}

	countDisliked, err := h.services.CountDislikedPosts(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}

	comment, err := h.services.GetCommentByPostId(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	for c := 0; c < len(comment); c++ { // here you can find how much like this post got
		count, err := h.services.CountLikedComment(comment[c].Comment_id) // by this function
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		comment[c].Likes = count
	}
	for d := 0; d < len(comment); d++ { // here you can find how much like this post got
		count, err := h.services.CountDislikedComment(comment[d].Comment_id) // by this function
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		comment[d].Dislikes = count
	}
	u := models.DisplayPost{
		Author:        post.Author,
		Title:         post.Title,
		Posts:         post.Posts,
		Countliked:    countLiked,
		Countdisliked: countDisliked,
		Forcomment:    comment,
	}
	if err = h.render(writer, "display.html", u); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
}
