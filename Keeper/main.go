// Copyright (c) 2018 SHIGEMUNE Katsuhiro
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
package main

//Module:000
import (
	"./Settings"
	"./Receiver"
	"./Verifier"
	"./Sender"
	"flag"
	"log"
	"math"
	"strconv"
	"fmt"
	"time"
)

var (
	ChangeOpt  = flag.Int("c", -1, "cオプション:送信間隔変更[定義域:1～6] 例) -c 1")
	LateOpt  = flag.Bool("l",false, "lオプション:遅延補正[補正時間はメールファイルから計算する。] 例) -l")
	LateOnValueOpt  = flag.Int("lx",math.MaxInt32, "lxオプション:遅延補正[定義域:-1800～1800秒] 例) -lx -300")
	LateOnPipeOpt = flag.Bool("lp", false, "lpオプション:遅延補正[最新の受信日時をパイプで受け取る。フォーマット:yyyy-MM-ddTHH:mm:ss(ISO8601)] 例) -lp")
	MailAddressOpt = flag.String("m", "", "mオプション:送信先メールアドレス 例) -\"m\" test@test.com")
	IdOpt  = flag.String("i", "", "iオプション:任意の識別子[定義域:半角アルファベット16文字] 例) -i \"SV BUOY\"")
	Offset1Opt  = flag.Float64("1", math.MaxFloat64, "1オプション:1項目目の補正値[定義域:-1.0～1.0] 例) -1 -1.0\n注:オフセットして下げた電圧値が電圧制御のしきい値を下回ると動作停止します。(=遠隔では制御不能になります。)")
	Offset2Opt  = flag.Float64("2", math.MaxFloat64, "2オプション:2項目目の補正値[定義域:-5.0～5.0] 例) -2 -3.5")
	Offset3Opt  = flag.Float64("3", math.MaxFloat64, "3オプション:3項目目の補正値[定義域:-5.0～5.0] 例) -3 -3.5")
	Offset4Opt  = flag.Float64("4", math.MaxFloat64, "4オプション:4項目目の補正値[定義域:-5.0～5.0] 例) -4 -3.5")
	VoltOpt  = flag.String("v", "", "vオプション:電圧制御[定義域:ON or OFF] 例) -v ON")
	HalfOpt  = flag.String("h", "", "hオプション:送信回数[定義域:ON or OFF] 例) -h ON")
	XOpt  = flag.Float64("extension", math.MaxFloat64, "xオプション:動作停止電圧[定義域:7.0～12.0] 例) -x 10.5\n注:電圧制御のしきい値が現在の電圧を下回ると動作停止します。(=遠隔では制御不能になります。)")
	TestOpt  = flag.String("optional", "", "optionalオプション：任意の設定データを送信する。: 例) -optional")
	TimeOpt  = flag.Bool("time", false, "timeオプション：システムから正常な時間が得られるかを確認する。: 例) -time")
)

func main() {

	Settings.ReadSettingsFile()

	flag.Parse()
	if -1 != *ChangeOpt {
		if Verifier.IsSendingPeriod(*ChangeOpt) {
			log.Printf(",001,送信間隔変更の設定が実行されました。（設定モード：%d）\n", *ChangeOpt)
			Sender.SendStringByMail("$C,"+strconv.Itoa(*ChangeOpt), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",002,送信間隔変更の設定が実行できませんでした。")
		}
	} else if *LateOpt {
		NewDT,OldDT,hasErr := Receiver.GetRecentMailDateTime(Settings.SettingsXml.Config.MailtextPath)
		if hasErr {
			log.Println(",004,遅延補正が実行できませんでした。")
			return
		}
		isAction, AdjustmentSec := Verifier.GetSettingsSec(NewDT,OldDT,Settings.SettingsXml.Config)
		if isAction {
			log.Printf(",003,遅延補正が実行されました。（設定秒数：%d秒）\n", AdjustmentSec)
			Sender.SendStringByMail("$L,"+strconv.Itoa(AdjustmentSec), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",004,遅延補正が実行できませんでした。")
		}
	} else if math.MaxInt32 != *LateOnValueOpt {
		if Verifier.IsLateValue(*LateOnValueOpt) {
			log.Printf(",003,遅延補正が実行されました。（設定秒数：%d秒）\n", *LateOnValueOpt)
			Sender.SendStringByMail("$L,"+strconv.Itoa(*LateOnValueOpt), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",004,遅延補正が実行できませんでした。")
		}
	}else if *LateOnPipeOpt{
		var stdin string
		fmt.Scan(&stdin)
		//
		old,new,isOK := Verifier.IsLateDateTime(stdin)
		if isOK {
			isAction, AdjustmentSec := Verifier.GetSettingsSec(old,new, Settings.SettingsXml.Config)
			if isAction {
				log.Printf(",003,遅延補正が実行されました。（設定秒数：%d秒）\n", AdjustmentSec)
				Sender.SendStringByMail("$L,"+strconv.Itoa(AdjustmentSec), Settings.SettingsXml.Smtp)
			} else {
				log.Println(",004,遅延補正が実行できませんでした。")
			}
		} else {
			log.Println(",004,遅延補正が実行できませんでした。")
		}
	} else if "" != *MailAddressOpt {
		if Verifier.IsMailAddress(*MailAddressOpt) {
			log.Printf(",005,送信先メールアドレスの設定が実行されました。（設定アドレス：%s）\n", *MailAddressOpt)
			Sender.SendStringByMail("$M," + *MailAddressOpt, Settings.SettingsXml.Smtp)
		} else {
			log.Println(",006,送信先メールアドレスの設定が実行できませんでした。")
		}

	} else if "" != *IdOpt {
		if Verifier.IsID(*IdOpt) {
			log.Printf(",007,任意識別子の登録が実行されました。（設定ID：%s）\n", *IdOpt)
			Sender.SendStringByMail("$I," + *IdOpt, Settings.SettingsXml.Smtp)
		} else {
			log.Println(",008,任意識別子の登録が実行できませんでした。")
		}

	} else if math.MaxFloat64 != *Offset1Opt {
		if Verifier.IsOffset4Voltage(*Offset1Opt) {
			log.Printf(",009,電圧の計測値の補正値の登録が実行されました。（補正値：%f）\n", *Offset1Opt)
			Sender.SendStringByMail("$1,"+strconv.FormatFloat(*Offset1Opt, 'f', 1, 64), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",010,電圧の計測値の補正値の登録が実行されませんでした。(注：電圧の定義域は、-1.0～+1.0までです。)")
		}

	} else if math.MaxFloat64 != *Offset2Opt {
		if Verifier.IsOffset(*Offset2Opt) {
			log.Printf(",011,2列目の計測値の補正値の登録が実行されました。（補正値：%f）\n", *Offset2Opt)
			Sender.SendStringByMail("$2,"+strconv.FormatFloat(*Offset2Opt, 'f', 1, 64), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",012,2列目の計測値の補正値の登録が実行されませんでした。")
		}
	} else if math.MaxFloat64 != *Offset3Opt {
		if Verifier.IsOffset(*Offset3Opt) {
			log.Printf(",013,3列目の計測値の補正値の登録が実行されました。（補正値：%f）\n", *Offset3Opt)
			Sender.SendStringByMail("$3,"+strconv.FormatFloat(*Offset3Opt, 'f', 1, 64), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",014,3列目の計測値の補正値の登録が実行されませんでした。")
		}
	} else if math.MaxFloat64 != *Offset4Opt {
		if Verifier.IsOffset(*Offset4Opt) {
			log.Printf(",015,4列目の計測値の補正値の登録が実行されました。（補正値：%f）\n", *Offset4Opt)
			Sender.SendStringByMail("$4,"+strconv.FormatFloat(*Offset4Opt, 'f', 1, 64), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",016,4列目の計測値の補正値の登録が実行されませんでした。")
		}
	} else if "" != *VoltOpt {
		if "ON" == *VoltOpt || "on" == *VoltOpt {
			log.Println(",017,電圧制御ONが設定されました。")
			Sender.SendStringByMail("$V,ON", Settings.SettingsXml.Smtp)
		} else if "OFF" == *VoltOpt || "off" == *VoltOpt {
			log.Println(",018,電圧制御OFFが設定されました。")
			Sender.SendStringByMail("$V,OFF", Settings.SettingsXml.Smtp)
		} else {
			log.Println(",019,電圧制御が設定できませんでした。")
		}
	} else if "" != *HalfOpt {
		if "ON" == *HalfOpt || "on" == *HalfOpt {
			log.Println(",020,送信回数ON[30分に1回]が設定されました。")
			Sender.SendStringByMail("$H,ON", Settings.SettingsXml.Smtp)
		} else if "OFF" == *HalfOpt || "off" == *HalfOpt {
			log.Println(",021,送信回数OFF[1時間に1回]が設定されました。")
			Sender.SendStringByMail("$H,OFF", Settings.SettingsXml.Smtp)
		} else {
			log.Println(",022,送信回数の設定ができませんでした。")
		}

	} else if math.MaxFloat64 != *XOpt {
		if Verifier.IsTerminationVoltage(*XOpt) {
			log.Printf(",023,動作停止電圧値が設定されました。（設定値：%f）\n", *XOpt)
			Sender.SendStringByMail("$X,"+strconv.FormatFloat(*XOpt, 'f', 1, 64), Settings.SettingsXml.Smtp)
		} else {
			log.Println(",024,動作停止電圧値が設定できませんでした。")
		}
	} else if "" != *TestOpt {
		log.Println(",025,テスト用に入力したデータの送信が実行されました。",)
		Sender.SendStringByMail(*TestOpt, Settings.SettingsXml.Smtp)
	}else if *TimeOpt {
		log.Printf(",026,現在のシステム日時は、%s です。",time.Now().String())
	}
}
