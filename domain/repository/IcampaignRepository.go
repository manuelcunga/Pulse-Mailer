package repository

import entities "github.com/manuelcunga/Pulse-Mailer/domain/entities/campaign"

type IcampaignRepository interface {
	Create(campaign *entities.Campaign) error
	Update(campaign *entities.Campaign) error
	FindAll() ([]*entities.Campaign, error)
	GetByID(id string) (*entities.Campaign, error)
	Delete(id string) error
}
