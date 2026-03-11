package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/MorningStar264/Url_shortner/internal/helper"
	"github.com/MorningStar264/Url_shortner/internal/model"
	"github.com/MorningStar264/Url_shortner/internal/repository"
	"github.com/MorningStar264/Url_shortner/internal/server"
	"github.com/go-chi/chi"
)

type LinkHandler struct {
	server       *server.Server
	repositories *repository.Repositories
}

func NewLinkHandler(srv *server.Server, repo *repository.Repositories) *LinkHandler {
	return &LinkHandler{
		server:       srv,
		repositories: repo,
	}
}
func (h *LinkHandler) GenerateShortenedUrl(w http.ResponseWriter, r *http.Request) {
	var l model.Link

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// ensure the unknown field are not there
	decoder.DisallowUnknownFields()
	// decoding the data
	err := decoder.Decode(&l)

	l.ShortCode = h.server.Node.GenerateID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repositories.LinkMethods.CreateLink(l)
	fmt.Println("generate url")
	fmt.Println(l)
	w.Write([]byte(l.ShortCode))
}
func (h *LinkHandler) RedirectUrl(w http.ResponseWriter, r *http.Request) {
	short_url := chi.URLParam(r, "shortened_Url")

	var l model.Link
	l.ShortCode = short_url

	link := h.repositories.LinkMethods.GetLink(l)
	fmt.Println("redirecting to ", link.LongURL)
	http.Redirect(w, r, link.LongURL, http.StatusFound)

}
