package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Contact struct {
	Email string `json:"email" valid:"email,required~ O campo E-mail é obrigatório."`
}

type Campaign struct {
	ID        string    `json:"id" valid:"uuid,required~ O ID da campanha é obrigatório."`
	Name      string    `json:"name" valid:"required~ O campo Nome é obrigatório."`
	Content   string    `json:"content" valid:"required~ O campo conteúdo é obrigatório."`
	Contacts  []Contact `json:"contacts" valid:"required~ O campo E-mails é obrigatório."`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, value := range emails {
		contacts[index].Email = value
	}

	campaign := Campaign{
		ID:        uuid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}

	if err := campaign.isValid(); err != nil {
		return nil, err
	}

	return &campaign, nil
}

func (cp *Campaign) isValid() error {
	_, err := govalidator.ValidateStruct(cp)
	if err != nil {
		return err
	}
	return nil
}
