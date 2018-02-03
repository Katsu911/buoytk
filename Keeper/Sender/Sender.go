package Sender

import (
	"../Settings"
	"net/smtp"
	"strconv"
)

func SendSettingMin(min int, server Settings.Smtp) bool {

	auth := smtp.PlainAuth(
		"",
		server.Account,
		server.Password,
		server.Host,
	)

	if err := smtp.SendMail(server.Host, auth, server.Source, []string{server.Address}, []byte("$L," + strconv.Itoa(min))); err != nil {
		return false
	}
	return true
}
