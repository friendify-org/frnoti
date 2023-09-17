package mail_route

import (
	"main/config"
	"main/dtos"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

func New(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.Email{})

	emailService := services.EmailService{}

	route = route.Group("/mails") // Send mail
	{
		route.POST("", middlewares.BindBody[config.MailData]("data"), func(context *gin.Context) {
			emailService.Send(context)
		})
		route.GET("", middlewares.BindQuery[dtos.Query]("query"), emailService.GetAll)
	}
}
