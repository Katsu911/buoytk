package Sender

import (
	"../Settings"
	"net/smtp"
	"log"
	"os"
	"time"
	"strconv"
)

func logingSendMail(f string){
	file, err := os.Create(f)
	if err != nil {
		log.Printf(",200,メール送信履歴を書きこめませんでした。（ファイル名:%s）\n", f)
	}
	defer file.Close()

	tm:=time.Now()
	ye, mo, da := tm.Date()
	output := strconv.Itoa(ye) + "," +
		strconv.Itoa((int)(mo)) + "," +
		strconv.Itoa(da) + "," +
	strconv.Itoa(tm.Hour()) + "," +
	strconv.Itoa(tm.Minute()) + "," +
	strconv.Itoa(tm.Second())

	file.Write(([]byte)(output))
}

func SendStringByMail(mes string, server Settings.Smtp) {

	auth := smtp.PlainAuth(
		"",
		server.Account,
		server.Password,
		server.Host,
	)

	err := smtp.SendMail(
		server.Host+":587",
		auth,
		server.Source,
		[]string{server.Address},
		[]byte(mes),
	)
	if err != nil {
		log.Printf(",101,メール送信に失敗しました（Host:%s To:%s  Body:%s）\n", server.Host, server.Address, mes)
		return
	}
	logingSendMail("SendingHistoryOfSettingMail")
	log.Printf(",100,メールを送信しました。（Host:%s To:%s  Body:%s）\n", server.Host, server.Address, mes)
	return
}