package campaign_usecase

import (
	entities "github.com/manuelcunga/Pulse-Mailer/domain/entities/campaign"
	"github.com/manuelcunga/Pulse-Mailer/domain/repository"
	"github.com/manuelcunga/Pulse-Mailer/usecases/campaign/dtos"
)

type CampaignUsecase interface {
	Create(campaign *dtos.CreateCampaignInput) (string, error)
}

type CampaignUseCaseImplment struct {
	CampaignRepo repository.IcampaignRepository
}

func NewCampaignUseCase(campaignRepo repository.IcampaignRepository) CampaignUsecase {
	return &CampaignUseCaseImplment{
		CampaignRepo: campaignRepo,
	}
}

func (cp CampaignUseCaseImplment) Create(data *dtos.CreateCampaignInput) (string, error) {
	campaign, err := entities.NewCampaign(data.Name, data.Content, data.Contacts)

	if err != nil {
		return "", err
	}

	err = cp.CampaignRepo.Create(campaign)

	if err != nil {
		return "", err
	}
	return campaign.ID, nil

}
