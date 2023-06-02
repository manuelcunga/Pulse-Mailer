package repository

import (
	entities "github.com/manuelcunga/Pulse-Mailer/domain/entities/campaign"
	"github.com/manuelcunga/Pulse-Mailer/domain/repository"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IcampaignRepository {
	return &CampaignRepository{Db: db}
}

func (u *CampaignRepository) Create(campaign *entities.Campaign) error {
	return u.Db.Create(campaign).Error
}

func (u *CampaignRepository) Delete(id string) error {
	return u.Db.Delete(&entities.Campaign{}, id).Error
}

func (u *CampaignRepository) FindAll() ([]*entities.Campaign, error) {
	var campaign []*entities.Campaign
	err := u.Db.Find(&campaign).Error
	return campaign, err
}

func (u *CampaignRepository) GetByID(id string) (*entities.Campaign, error) {
	var campaign entities.Campaign
	err := u.Db.First(&campaign, id).Error
	return &campaign, err
}

func (u *CampaignRepository) Update(campaign *entities.Campaign) error {
	return u.Db.Save(campaign).Error
}
