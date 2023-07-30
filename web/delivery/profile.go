// done clean architecture
package delivery

import (
	"errors"
	"fmt"
	"net/http"

	"forum/web/models"
)

func (h *Handler) profile(writer http.ResponseWriter, request *http.Request) { // methods include
	if request.URL.Path != "/profile" {
		h.errorPage(writer, request, http.StatusNotFound)
		return
	}

	userId := request.Context().Value(userIdCtx)
	if userId == 0 {
		h.errorPage(writer, request, http.StatusUnauthorized)
		return
	}
	username, email, err := h.services.GetUserService(userId.(int))
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}

	allPost, err := h.services.GetAllPostStorage()
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	u := models.User{
		Username: username,
		Email:    email,
		Post:     allPost, // an array to create all the posts, author, title,content
	}

	switch request.Method {
	case http.MethodGet:
		if err = h.render(writer, "profile.html", u); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		userId := request.Context().Value(userIdCtx)
		if err := request.ParseForm(); err != nil {
			fmt.Println("err:", err)
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}

		t, ok := request.Form["title"]
		if !ok {
			u.Erstring = "Incorrect title"
			writer.WriteHeader(http.StatusBadRequest)
			if err := h.render(writer, "profile.html", u); err != nil {
				h.errorPage(writer, request, http.StatusInternalServerError)
			}
			return
		}
		title := t[0]

		p, ok := request.Form["mypost"]
		if !ok {
			u.Erstring = "Incorrect post"
			writer.WriteHeader(http.StatusBadRequest)
			if err := h.render(writer, "profile.html", u); err != nil {
				h.errorPage(writer, request, http.StatusInternalServerError)
			}
			return
		}
		post := p[0]
		fmt.Println("post: ", post)
		// for _, el := range post {
		// 	if el < 32 || el > 126 {
		// 		fmt.Println("non-ascii")
		// 		u.Erstring = "Non latin alphabet is not supported"
		// 		if err := h.render(writer, "profile.html", u); err != nil {
		// 			h.errorPage(writer, request, http.StatusInternalServerError)
		// 		}
		// 		return
		// 	}
		// }

		contents, ok := request.Form["content"] //[sales,holilday,books]
		if !ok {
			u.Erstring = "Please choose at least one category"
			writer.WriteHeader(http.StatusBadRequest)
			if err := h.render(writer, "profile.html", u); err != nil {
				h.errorPage(writer, request, http.StatusInternalServerError)
			}
			return
		}
		if err := CheckContents(contents); err != nil {
			fmt.Println("CheckContents error")
			h.errorPage(writer, request, http.StatusBadRequest)
			return
		}
		username, _, err := h.services.GetUserById(userId.(int))
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		if IsTagEmpty(title) { // need to finish this one, but to work with post method
			u.Erstring = "Title cannot be empty"
			writer.WriteHeader(http.StatusBadRequest)
			if err := h.render(writer, "profile.html", u); err != nil {
				h.errorPage(writer, request, http.StatusInternalServerError)
			}
			return
		}

		if IsTagEmpty(post) { // need to finish this one, but to work with post method
			u.Erstring = "Post cannot be empty"
			writer.WriteHeader(http.StatusBadRequest)
			if err := h.render(writer, "profile.html", u); err != nil {
				h.errorPage(writer, request, http.StatusInternalServerError)
			}
			return
		}

		if err := h.services.FillThePostTable(userId.(int), post, username, title); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}

		post_id, err := h.services.FindPostIdbyPost(post)
		if err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		if err := h.services.FillTheContentTable(post_id, contents); err != nil {
			h.errorPage(writer, request, http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/profile", http.StatusMovedPermanently)
	}
}

func (h *Handler) profileContentDisplay(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		h.errorPage(writer, request, http.StatusMethodNotAllowed)
		return
	}
	userId := request.Context().Value(userIdCtx).(int)
	if userId == 0 {
		h.errorPage(writer, request, http.StatusUnauthorized)
		return
	}
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
	if err := CheckContents(b); err != nil {
		h.errorPage(writer, request, http.StatusBadRequest)
		return
	}
	buttons := b[0]
	displayallpost, err := h.services.GetAllPostsByCategory(buttons)
	if err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
	u := models.User{
		Username: username,
		Email:    email,
		Post:     displayallpost,
	}
	if err = h.render(writer, "profile.html", u); err != nil {
		h.errorPage(writer, request, http.StatusInternalServerError)
		return
	}
}

func CheckContents(contents []string) error {
	for _, each := range contents {
		if each != "sales" && each != "holiday" && each != "facts" && each != "events" && each != "quotes" && each != "books" && each != "guidesandtips" {
			return errors.New("does not match")
		}
	}
	return nil
}
