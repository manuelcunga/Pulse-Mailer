package campaign_usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manuelcunga/Pulse-Mailer/domain/repository/mocks"
	campaign_usecase "github.com/manuelcunga/Pulse-Mailer/usecases/campaign"
	"github.com/manuelcunga/Pulse-Mailer/usecases/campaign/dtos"
	"github.com/stretchr/testify/assert"
)

func TestCreateCampaign_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCampaignRepo := mocks.NewMockIcampaignRepository(ctrl)

	campaignUseCase := campaign_usecase.NewCampaignUseCase(mockCampaignRepo)

	input := &dtos.CreateCampaignInput{
		Name:     "Marketing",
		Content:  "Welcome to our app",
		Contacts: []string{"manuel@gmail.com", "android2@gmail.com"},
	}

	mockCampaignRepo.EXPECT().Create(gomock.Any()).Return(nil)

	campaignID, err := campaignUseCase.Create(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, campaignID)
}

func TestCreateCampaign_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCampaignRepo := mocks.NewMockIcampaignRepository(ctrl)

	campaignUseCase := campaign_usecase.NewCampaignUseCase(mockCampaignRepo)

	// Input with missing required fields
	invalidInput := &dtos.CreateCampaignInput{}

	campaignID, err := campaignUseCase.Create(invalidInput)

	assert.Error(t, err)
	assert.Empty(t, campaignID)
}

func TestCreateCampaign_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCampaignRepo := mocks.NewMockIcampaignRepository(ctrl)

	campaignUseCase := campaign_usecase.NewCampaignUseCase(mockCampaignRepo)

	input := &dtos.CreateCampaignInput{
		Name:     "Marketing",
		Content:  "Welcome to our app",
		Contacts: []string{"manuel@gmail.com", "android2@gmail.com"},
	}

	expectedError := errors.New("Internal Server Error")
	mockCampaignRepo.EXPECT().Create(gomock.Any()).Return(expectedError)

	campaignID, err := campaignUseCase.Create(input)

	assert.Equal(t, expectedError, err)
	assert.Empty(t, campaignID)
}
