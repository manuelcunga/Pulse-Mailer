package dtos

type CreateCampaignInput struct {
	Name     string   `json:"name" `
	Content  string   `json:"content"`
	Contacts []string `json:"contacts" `
}
