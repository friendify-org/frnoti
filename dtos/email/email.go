package email_dto

type SendMail struct {
	To      string `json:"to"`
	Content string `json:"content"`
	Subject string `json:"subjet"`
}
