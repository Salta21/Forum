// everything related to posts
package delivery

import (
	"net/http"
	"strconv"

	"forum/web/models"
)

func (h *Handler) onlyCreatedPosts(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/created-posts" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	username, email, err := h.services.GetUserById(userId)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if err := request.ParseForm(); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	p, ok := request.Form["buttons"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	buttons := p[0]
	if buttons == "createdposts" { // buttons for liked and disliked posts and create database
		allPosts, err := h.services.GetPostsById(userId)
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		u := models.User{
			Username: username,
			Email:    email,
			Post:     allPosts,
		}
		if err = h.render(writer, "profile.html", u); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	} else {
		h.errorPage(writer, request, http.StatusBadRequest)
	}
}

func (h *Handler) onlyLikedPosts(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/liked-posts" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	username, email, err := h.services.GetUserById(userId)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if err := request.ParseForm(); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	b, ok := request.Form["buttons"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	buttons := b[0]
	if buttons == "likedposts" { // buttons for liked and disliked posts and create database
		var allPostIdOfLikedPosts []models.LikedPosts
		var res []models.Post
		allPostIdOfLikedPosts, err = h.services.GetLikedPostsById(userId)
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		for _, each := range allPostIdOfLikedPosts { // thats why there are two pages
			hello, err := h.services.GetPostByPostId(each.Post_id)
			if err != nil {
				h.errorPage(writer, request, http.StatusInternalServerError)
				return
			}
			res = append(res, hello)
		}
		u := models.User{
			Username: username,
			Email:    email,
			Post:     res,
		}
		if err := h.render(writer, "profile.html", u); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	} else {
		h.errorPage(writer, request, http.StatusBadRequest)
	}
}

func (h *Handler) dislikePost(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/dislikepost" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if err := request.ParseForm(); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	p, ok := request.Form["post"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	post_id := p[0]
	u, ok := request.Form["user"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	user_id := u[0]
	postId, err := strconv.Atoi(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	usId, err := strconv.Atoi(user_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if usId != userId {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	checkforLike, err := h.services.IfUserLikedPost(user_id, post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if checkforLike {
		if err := h.services.DeleteLikedPost(usId, postId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	}
	bl, err := h.services.IfUserDislikedPost(user_id, post_id)
	if err == nil && bl {
		if err := h.services.DeleteDislikedPost(usId, postId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusMovedPermanently)
	}
	if !bl {
		if err := h.services.FillTheDislikesPostTable(usId, postId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusMovedPermanently)
	}
}

func (h *Handler) likePost(writer http.ResponseWriter, request *http.Request) { // здесь нужно проверить дислайкал ли человек, если да то удаляем с дислайка этот юзер и пост и только потом заполняем лайк таблицу
	if request.URL.Path != "/likepost" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if err := request.ParseForm(); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	p, ok := request.Form["post"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	post_id := p[0]
	u, ok := request.Form["user"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	user_id := u[0]
	postId, err := strconv.Atoi(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	usId, err := strconv.Atoi(user_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if userId != usId {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	checkforDislike, err := h.services.IfUserDislikedPost(user_id, post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}

	if checkforDislike {
		if err := h.services.DeleteDislikedPost(usId, postId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	}
	bl, err := h.services.IfUserLikedPost(user_id, post_id)
	if err == nil && bl {
		if err := h.services.DeleteLikedPost(usId, postId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusFound)
		return
	}
	if !bl {
		if err := h.services.FillTheLikesPostTable(usId, postId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusFound)
	}
}
