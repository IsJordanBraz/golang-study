package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	meuId := chi.URLParam(r, "date")
	println(w)
	println(meuId)
	println(r)
	campaign, err := h.CampaignService.GetBy(id)
	return campaign, 200, err
}
