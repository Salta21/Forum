package delivery

import (
	"net/http"
	"strconv"

	"forum/web/models"
)

func (h *Handler) moreInfo(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/moreinfo/" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	user_id := request.Context().Value(userIdCtx)
	post_id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	allPost, err := h.services.GetPostByPostId(post_id)
	if err != nil {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	countLike, err1 := h.services.CountLikedPosts(post_id)
	if err1 != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}

	countDislike, err2 := h.services.CountDislikedPosts(post_id)
	if err2 != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	comments, err := h.services.GetCommentByPostId(post_id)
	if err != nil {
		type display struct {
			P            models.Post
			U            int
			CountLike    int
			CountDislike int
		}
		info := display{P: allPost, U: user_id.(int), CountLike: countLike, CountDislike: countDislike}
		if err = h.render(writer, "moreinfo.html", info); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	}
	type display struct {
		P            models.Post
		U            int
		CountLike    int
		CountDislike int
		Comments     []models.PostTheComment
	}

	for c := 0; c < len(comments); c++ { // here you can find how much like this post got
		count, err := h.services.CountLikedComment(comments[c].Comment_id) // by this function
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		comments[c].Likes = count
	}
	for d := 0; d < len(comments); d++ { // here you can find how much like this post got
		count, err := h.services.CountDislikedComment(comments[d].Comment_id) // by this function
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		comments[d].Dislikes = count
	}
	info := display{P: allPost, U: user_id.(int), CountLike: countLike, CountDislike: countDislike, Comments: comments}
	if err = h.render(writer, "moreinfo.html", info); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
}
