package campaign_controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	campaign_controllers "github.com/manuelcunga/Pulse-Mailer/api/rest/Campaign/controllers/create"
	"github.com/manuelcunga/Pulse-Mailer/domain/repository/mocks"
	campaign_usecase "github.com/manuelcunga/Pulse-Mailer/usecases/campaign"
	"github.com/manuelcunga/Pulse-Mailer/usecases/campaign/dtos"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCampaign(t *testing.T) {
	t.Run("Should be create a new Campaign", func(t *testing.T) {
		e := echo.New()

		input := &dtos.CreateCampaignInput{
			Name:     "Marketing",
			Content:  "Welcome to our app",
			Contacts: []string{"manuel@gmail.com", "android2@gmail.com"},
		}

		jsonData, _ := json.Marshal(input)

		res := httptest.NewRequest("POST", "/api/v1/campaign/create", bytes.NewBuffer(jsonData))
		res.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ctx := e.NewContext(res, rec)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCampaignRepository := mocks.NewMockIcampaignRepository(ctrl)
		mockCampaignRepository.EXPECT().Create(gomock.Any()).Return(nil)

		campaignUseCase := campaign_usecase.NewCampaignUseCase(mockCampaignRepository)
		campaignController := campaign_controllers.NewCampaignController(campaignUseCase)

		if assert.NoError(t, campaignController.Handle(ctx)) {
			assert.Equal(t, http.StatusCreated, rec.Code, "Error status code != 201")
		}

	})
}
