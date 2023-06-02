package campaign_controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	erro "github.com/manuelcunga/Pulse-Mailer/api/rest/errors"
	campaign_usecase "github.com/manuelcunga/Pulse-Mailer/usecases/campaign"
	"github.com/manuelcunga/Pulse-Mailer/usecases/campaign/dtos"
)

type CampaignController struct {
	campaignUsecase campaign_usecase.CampaignUsecase
}

func NewCampaignController(campaignUse campaign_usecase.CampaignUsecase) *CampaignController {
	return &CampaignController{campaignUsecase: campaignUse}
}

func (campaignCtl *CampaignController) Handle(ctx echo.Context) error {
	var campaign dtos.CreateCampaignInput

	if err := ctx.Bind(&campaign); err != nil {
		return ctx.JSON(http.StatusBadRequest, erro.ErrorResponse{Message: err.Error()})
	}

	output, err := campaignCtl.campaignUsecase.Create(&campaign)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, erro.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, output)
}
