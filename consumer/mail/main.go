package mail_consumer

import (
	"encoding/json"
	"fmt"
	base_consumer "main/consumer"
	"main/models"
	mail_repository "main/repositories/email"
)

func New() {
	event := base_consumer.New("mail")
	event.SubcribeEvent(sendMailListener)
}

func sendMailListener(recieved []byte) {
	var data models.Email
	err := json.Unmarshal(recieved, &data)
	if err != nil {
		fmt.Printf("error when decode json data is: %v\n", err)
		return
	}
	if err := mail_repository.SendMailToUser(&data); err != nil {
		fmt.Printf("Error when send mail is: %v %v\n", err, data)
	}
}
