package delivery

import (
	"net/http"

	"forum/web/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.middleWare(h.home))
	mux.HandleFunc("/displayforunauthorized/", h.middleWare(h.displayForUnauthorized))
	mux.HandleFunc("/sign-out", h.middleWare(h.signOut))
	mux.HandleFunc("/commentpost", h.middleWare(h.commentPost))
	mux.HandleFunc("/likepost", h.middleWare(h.likePost))
	mux.HandleFunc("/dislikepost", h.middleWare(h.dislikePost))
	mux.HandleFunc("/profilecontentdisplay", h.middleWare(h.profileContentDisplay))
	mux.HandleFunc("/created-posts", h.middleWare(h.onlyCreatedPosts))
	mux.HandleFunc("/liked-posts", h.middleWare(h.onlyLikedPosts))
	mux.HandleFunc("/dislikecomment", h.middleWare(h.dislikeComment))
	mux.HandleFunc("/likecomment", h.middleWare(h.likeComment))
	mux.HandleFunc("/signup", h.middleWare(h.signUp))
	mux.HandleFunc("/moreinfo/", h.middleWare(h.moreInfo))
	mux.HandleFunc("/signin", h.middleWare(h.signIn))
	mux.HandleFunc("/profile", h.middleWare(h.profile))

	fileServer := http.FileServer(http.Dir("./web/static/style/"))
	mux.Handle("/style/", http.StripPrefix("/style", fileServer))

	fileServer2 := http.FileServer(http.Dir("./web/static/img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fileServer2))
	return mux
}
