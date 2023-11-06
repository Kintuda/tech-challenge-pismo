package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kintuda/tech-challenge-pismo/pkg/account"
	"github.com/rs/zerolog/log"
)

type AccountController struct {
	service *account.AccountService
}

func NewAccountController(service *account.AccountService) *AccountController {
	return &AccountController{
		service: service,
	}
}

type CreateAccount struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

func (a *AccountController) CreateAccount(c *gin.Context) {
	payload := new(CreateAccount)

	if err := c.ShouldBindBodyWith(payload, binding.JSON); err != nil {
		if err := c.Error(err); err != nil {
			log.Error().Err(err).Msgf("error while handling error")
		}

		c.Abort()
		return
	}

	account, err := a.service.CreateAccount(c.Request.Context(), payload.DocumentNumber)

	if err != nil {
		if err := c.Error(err); err != nil {
			log.Err(err).Msgf("error while handling error %v", err)
		}
		c.Abort()
		return
	}

	c.JSON(201, account)
}

func (a *AccountController) GetAccount(c *gin.Context) {
	id := c.Param("id")

	account, err := a.service.GetAccount(c.Request.Context(), id)

	if err != nil {
		if err := c.Error(err); err != nil {
			log.Error().Err(err).Msg("error while handling error")
		}

		c.Abort()
		return
	}

	if account == nil {
		c.JSON(404, gin.H{"message": "account not found"})
		return
	}

	c.JSON(200, account)
}
