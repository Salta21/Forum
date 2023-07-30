package delivery

import (
	"net/http"
	"strconv"
)

func (h *Handler) commentPost(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/commentpost" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	if err := request.ParseForm(); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	p, ok := request.Form["post"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	post_id, err := strconv.Atoi(p[0])
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	c, ok := request.Form["user"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	commenter_id, err := strconv.Atoi(c[0])
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if commenter_id != userId {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	if err = h.services.GetPostbyPost_id(post_id); err != nil {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}

	ct, ok := request.Form["commenttext"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	commenttext := ct[0]
	if IsTagEmpty(commenttext) {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	if err := h.services.FillTheCommentTable(commenter_id, post_id, commenttext); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/moreinfo/?id="+p[0], http.StatusMovedPermanently)
}

func (h *Handler) dislikeComment(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	disliker_id := request.Context().Value(userIdCtx).(int)
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
	cId, ok := request.Form["comment"]
	if !ok {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	commentId, err := strconv.Atoi(cId[0]) // commentid integer
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	checkforLike, err := h.services.IfLikerLikedComment(disliker_id, commentId)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if checkforLike {
		if err := h.services.DeleteLikeComments(disliker_id, commentId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	}
	result, err := h.services.IfDislikerDislikedComment(disliker_id, commentId)
	if err == nil && result {
		if err := h.services.DeleteDislikeComments(disliker_id, commentId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusMovedPermanently)
	}
	if !result {
		if err := h.services.FillTheDislikeCommentsTable(disliker_id, commentId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusMovedPermanently)
	}
}

func (h *Handler) likeComment(writer http.ResponseWriter, request *http.Request) { // here you can use liker_id and comment_id
	if request.URL.Path != "/likecomment" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	liker_id := request.Context().Value(userIdCtx).(int)
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
	comment_id := request.FormValue("comment")
	commentId, err := strconv.Atoi(comment_id) // commentid integer
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	// checkiflikerdislikedcomment
	checkforDislike, err := h.services.IfDislikerDislikedComment(liker_id, commentId)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	if checkforDislike {
		if err := h.services.DeleteDislikeComments(liker_id, commentId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	}
	bl, err := h.services.IfLikerLikedComment(liker_id, commentId)
	if err == nil && bl {
		if err := h.services.DeleteLikeComments(liker_id, commentId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusMovedPermanently)
	}
	if !bl {
		if err := h.services.FillTheLikeCommentsTable(liker_id, commentId); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/moreinfo/?id="+post_id, http.StatusMovedPermanently)
	}
}
