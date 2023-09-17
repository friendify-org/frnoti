package main

import (
	"fmt"
	"main/config"
	mail_consumer "main/consumer/mail"
	mail_route "main/routes/mails"
	"main/validator"

	"github.com/gin-gonic/gin"
)

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

func main() {

	config.SetupEnvirontment()

	fmt.Printf("------------------- Connecting to database ----------------")

	config.ConnectDataBase()

	config.ConnectRedis()

	config.InitializeMailSender()

	config.ConnectQueue()

	r := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api/v1"

	validator.ValidatorBinding()

	api := r.Group("/api/v1")
	{
		mail_route.New(api)
	}

	mail_consumer.New()

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	r.Run()
}
