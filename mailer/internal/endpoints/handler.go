package endpoints

import "gomailer/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
