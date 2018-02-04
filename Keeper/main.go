package main

import (
	"./Settings"
	"./Receiver"
	"./Analyzer"
	"./Sender"
	"flag"
	"log"
	"math"
	"strconv"
)

var (
	ChangeOpt  = flag.Int("c", 0, "help message for \"c\" option")
	LateOpt  = flag.Bool("l",false, "help message for \"l\" option")
	MailAddressOpt = flag.String("m", "", "help message for \"m\" option")
	IdOpt  = flag.String("i", "", "help message for \"i\" option")
	Offset1Opt  = flag.Float64("1", math.MaxFloat64, "help message for \"1\" option")
	Offset2Opt  = flag.Float64("2", math.MaxFloat64, "help message for \"2\" option")
	Offset3Opt  = flag.Float64("3", math.MaxFloat64, "help message for \"3\" option")
	Offset4Opt  = flag.Float64("4", math.MaxFloat64, "help message for \"4\" option")
	VoltOpt  = flag.String("v", "", "help message for \"v\" option")
	HalfOpt  = flag.String("h", "", "help message for \"h\" option")
	XOpt  = flag.Float64("x", math.MaxFloat64, "help message for \"x\" option")
)

func main() {

	Settings.ReadSettingsFile()

	flag.Parse()
	if 0 != *ChangeOpt {
		if Analyzer.IsSendingPeriod(*ChangeOpt){
			log.Printf(",001,送信間隔変更の設定が実行されました。（設定モード：%d）\n",*ChangeOpt)
			Sender.SendStringByMail("$C,"+ strconv.Itoa(*ChangeOpt),Settings.SettingsXml.Smtp)
		}else{
			log.Println(",002,送信間隔変更の設定が実行できませんでした。")
		}
	} else if *LateOpt{
		RecentMailDateTime := Receiver.GetRecentMailDateTime(Settings.SettingsXml.Config.MailtextPath)
		isAction, AdjustmentSec := Analyzer.GetSettingsSec(RecentMailDateTime,Settings.SettingsXml.Config)
		if isAction {
			log.Printf(",003,遅延補正を実行されました。（設定秒数：%d秒）\n", AdjustmentSec)
			Sender.SendStringByMail("$L,"+ strconv.Itoa(AdjustmentSec), Settings.SettingsXml.Smtp)
		}else{
			log.Println(",004,遅延補正が実行できませんでした。")
		}

	}else if "" != *MailAddressOpt{
		if Analyzer.IsMailAddress(*MailAddressOpt){
			log.Printf(",005,送信先メールアドレスの設定が実行されました。（設定アドレス：%s）\n",*MailAddressOpt)
			Sender.SendStringByMail("$M,"+ *MailAddressOpt,Settings.SettingsXml.Smtp)
		}else{
			log.Println(",006,送信先メールアドレスの設定が実行できませんでした。")
		}

	}else if "" != *IdOpt{
		if Analyzer.IsID(*IdOpt){
			log.Printf(",007,任意識別子の登録が実行されました。（設定ID：%s）\n",*IdOpt)
			Sender.SendStringByMail("$I,"+ *IdOpt,Settings.SettingsXml.Smtp)
		}else{
			log.Println(",008,任意識別子の登録が実行できませんでした。")
		}

	}else if math.MaxFloat64 != *Offset1Opt{
		if Analyzer.IsOffset(*Offset1Opt){
			log.Printf(",009,1列目の計測値の補正値の登録が実行されました。（補正値：%f）\n",*Offset1Opt)
			Sender.SendStringByMail("$1,"+ strconv.FormatFloat(*Offset1Opt,'f',1,64),Settings.SettingsXml.Smtp)
		}else{
			log.Println(",010,1列目の計測値の補正値の登録が実行さませんでした。")
		}

	}else if math.MaxFloat64 != *Offset2Opt{
		if Analyzer.IsOffset(*Offset2Opt){
			log.Printf(",011,2列目の計測値の補正値の登録が実行されました。（補正値：%f）\n",*Offset2Opt)
			Sender.SendStringByMail("$2,"+ strconv.FormatFloat(*Offset2Opt,'f',1,64),Settings.SettingsXml.Smtp)
		}else{
			log.Println(",012,2列目の計測値の補正値の登録が実行さませんでした。")
		}
	}else if math.MaxFloat64 != *Offset3Opt{
		if Analyzer.IsOffset(*Offset3Opt){
			log.Printf(",013,3列目の計測値の補正値の登録が実行されました。（補正値：%f）\n",*Offset3Opt)
			Sender.SendStringByMail("$3,"+ strconv.FormatFloat(*Offset3Opt,'f',1,64),Settings.SettingsXml.Smtp)
		}else{
			log.Println(",014,3列目の計測値の補正値の登録が実行さませんでした。")
		}
	}else if math.MaxFloat64 != *Offset4Opt{
		if Analyzer.IsOffset(*Offset4Opt){
			log.Printf(",015,4列目の計測値の補正値の登録が実行されました。（補正値：%f）\n",*Offset4Opt)
			Sender.SendStringByMail("$4,"+ strconv.FormatFloat(*Offset4Opt,'f',1,64),Settings.SettingsXml.Smtp)
		}else{
			log.Println(",016,4列目の計測値の補正値の登録が実行さませんでした。")
		}
	}else if "" != *VoltOpt {
		if "ON" == *VoltOpt || "on" == *VoltOpt {
			log.Println(",017,電圧制御ONが設定されました。")
			Sender.SendStringByMail("$V,ON", Settings.SettingsXml.Smtp)
		} else if  "OFF" == *VoltOpt || "off" == *VoltOpt{
			log.Println(",018,電圧制御OFFが設定されました。")
			Sender.SendStringByMail("$V,OFF", Settings.SettingsXml.Smtp)
		}else{
			log.Println(",019,電圧制御が設定できませんでした。")
		}
	}else if "" != *HalfOpt{
		if "ON" == *VoltOpt || "on" == *VoltOpt {
			log.Println(",020,送信回数ON[30分に1回]が設定されました。")
			Sender.SendStringByMail("$H,ON",Settings.SettingsXml.Smtp)
		}else if "OFF" == *VoltOpt || "off" == *VoltOpt {
			log.Println(",021,送信回数OFF[1時間に1回]が設定されました。")
			Sender.SendStringByMail("$H,OFF",Settings.SettingsXml.Smtp)
		}else{
			log.Println(",022,送信回数の設定ができませんでした。")
		}

	}else if math.MaxFloat64 != *XOpt{
		if Analyzer.IsTerminationVoltage(*XOpt){
			log.Printf(",023,動作停止電圧値が設定されました。（設定値：%f）\n",*XOpt)
			Sender.SendStringByMail("$X,"+ strconv.FormatFloat(*XOpt,'f',1,64),Settings.SettingsXml.Smtp)
		}else{
			log.Println(",024,動作停止電圧値が設定できませんでした。")
		}
	}

}
