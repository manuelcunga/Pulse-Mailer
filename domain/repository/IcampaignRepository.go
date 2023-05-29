package repository

import (
	entities "github.com/manuelcunga/Pulse-Mailer/domain/entities/campaign"
)

type IcampaignRepository interface {
	Create(campaign *entities.Campaign) error
}
