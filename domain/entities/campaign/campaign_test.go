package entities_test

import (
	"testing"
	"time"

	entities "github.com/manuelcunga/Pulse-Mailer/domain/entities/campaign"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Marketing"
	content  = "body"
	contacts = []string{"manuel@gmail.com", "jacob@gmail.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := entities.NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNew_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := entities.NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

func TestNew_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := entities.NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedAt, now)

}
