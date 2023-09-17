package mail_repository

import (
	"fmt"
	"main/config"
	"main/models"
)

func SendMailToUser(data *models.Email) error {

	result := config.Db.Create(&data)
	if result.Error != nil {
		fmt.Printf("Failed to save email: %v\n", result.Error.Error())
	}
	fmt.Printf("result: %v\n", result)

	return config.MailSender.SendMail(config.MailData{
		To:      data.SendTo,
		Subject: data.Subject,
		Content: data.Content,
	})
}
