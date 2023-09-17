package services

import (
	"main/base"
	"main/config"
	"main/dtos"
	"main/models"
	mail_repository "main/repositories/email"

	"github.com/gin-gonic/gin"
)

type EmailService struct {
	base.Service[models.Email]
}

func (emailService EmailService) Send(context *gin.Context) {
	mailData := context.MustGet("data").(models.Email)

	err := mail_repository.SendMailToUser(&mailData)
	if err != nil {
		response := config.Response{
			Data: config.NoData(),
			Messages: []config.Message{
				{
					Description: err.Error(),
				},
			},
		}

		response.BadRequest(context)
		return
	}

	response := config.Response{
		Data: config.ResponseData{
			Result: mailData,
		},
		Messages: []config.Message{
			config.Messages["send_mail_success"],
		},
	}

	response.Created(context)
}

func (emailService EmailService) GetAll(context *gin.Context) {
	query := context.MustGet("query").(dtos.Query)

	result, total := emailService.Repository.FindAll(&models.Email{}, query)

	response := config.Response{
		Data: config.ResponseData{
			Result: result,
			Total:  total,
			Limit:  query.Limit,
			Offset: query.Offset,
		},
		Messages: []config.Message{config.Messages["get_success"]},
	}

	response.GetSuccess(context)
}
