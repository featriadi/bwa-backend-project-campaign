package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type DetailCampaignFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	BackerCount      int                      `json:"backer_count"`
	UserId           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatDetailCampaign(campaign Campaign) DetailCampaignFormatter {
	detailCampaignFormatter := DetailCampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		BackerCount:      campaign.BackerCount,
		UserId:           campaign.UserID,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		detailCampaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{
		Name:     user.Name,
		ImageURL: user.Avatar,
	}

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	detailCampaignFormatter.Perks = perks
	detailCampaignFormatter.User = campaignUserFormatter
	detailCampaignFormatter.Images = images

	return detailCampaignFormatter
}
