package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaign, err := h.CampaignService.GetAll()
	return campaign, 200, err
}
