package Sender
//Module:100
import (
	"../Settings"
	"net/smtp"
	"log"
	"os"
	"time"
	"strconv"
	"bufio"
)

func writeSendingRecord(f string){

	fp , err := os.Create(f) //上書き保存
	if err != nil {
		log.Printf(",102,メール送信履歴を書きこめませんでした。（ファイル名:%s）\n", f)
	}
	defer fp.Close()

	writer :=bufio.NewWriter(fp)
	tm:=time.Now()
	ye, mo, da := tm.Date()
	output := strconv.Itoa(ye) + "," +
		strconv.Itoa((int)(mo)) + "," +
		strconv.Itoa(da) + "," +
	strconv.Itoa(tm.Hour()) + "," +
	strconv.Itoa(tm.Minute()) + "," +
	strconv.Itoa(tm.Second())

	_, err = writer.WriteString(output)
	if err != nil {
		log.Printf(",102,メール送信履歴を書きこめませんでした。（ファイル名:%s）\n", f)
	}
	writer.Flush()
}

func SendStringByMail(mes string, server Settings.Smtp) {

	auth := smtp.PlainAuth(
		"",
		server.Account,
		server.Password,
		server.Host,
	)

	const MIME_TYPE="MIME-Version: 1.0\r\n"
	const CONTENT_TYPE="Content-Type: text/plain; charset=\"iso-2022-jp\"\r\n"
	const TRANSFER="Content-Transfer-Encoding: 7bit\r\n"

	err := smtp.SendMail(
		server.Host+":587",
		auth,
		server.Source,
		[]string{server.Address},
		[]byte( MIME_TYPE + CONTENT_TYPE + TRANSFER + mes),
	)
	if err != nil {
		log.Printf(",101,メール送信に失敗しました（Host:%s To:%s  Body:%s）\n", server.Host, server.Address, mes)
		return
	}
	writeSendingRecord("SendingHistoryOfSettingMail")
	log.Printf(",100,メールを送信しました。（Host:%s To:%s  Body:%s）\n", server.Host, server.Address, mes)
	return
}