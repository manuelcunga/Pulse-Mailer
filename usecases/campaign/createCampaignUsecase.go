package campaign_usecase

import (
	internal_errors "github.com/manuelcunga/Pulse-Mailer/api/rest/errors"
	entities "github.com/manuelcunga/Pulse-Mailer/domain/entities/campaign"
	"github.com/manuelcunga/Pulse-Mailer/domain/repository"
	"github.com/manuelcunga/Pulse-Mailer/usecases/campaign/dtos"
)

type CampaignUsecase interface {
	Create(campaign *dtos.CreateCampaignInput) (string, error)
}

type CampaignUseCaseRepo struct {
	CampaignRepo repository.IcampaignRepository
}

func NewCampaignUseCase(campaignRepo repository.IcampaignRepository) CampaignUsecase {
	return &CampaignUseCaseRepo{
		CampaignRepo: campaignRepo,
	}
}

func (cp CampaignUseCaseRepo) Create(data *dtos.CreateCampaignInput) (string, error) {
	campaign, err := entities.NewCampaign(data.Name, data.Content, data.Contacts)

	if err != nil {
		return "", err
	}

	err = cp.CampaignRepo.Create(campaign)

	if err != nil {
		return "", internal_errors.ErrInternal
	}
	return campaign.ID, nil

}
