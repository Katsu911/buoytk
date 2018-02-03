package main

import (
	"./Settings"
	"./Receiver"
	"./Analyzer"
	"./Sender"
	"flag"
	"log"
)

var (
	ChangeOpt  = flag.Bool("c", false, "help message for \"c\" option")
	LateOpt  = flag.Bool("l",false, "help message for \"l\" option")
	MailAddressOpt = flag.Bool("m", false, "help message for \"m\" option")
	IdOpt  = flag.Bool("i", false, "help message for \"i\" option")
	Offset1Opt  = flag.Bool("1", false, "help message for \"1\" option")
	Offset2Opt  = flag.Bool("2", false, "help message for \"2\" option")
	Offset3Opt  = flag.Bool("3", false, "help message for \"3\" option")
	Offset4Opt  = flag.Bool("4", false, "help message for \"4\" option")
	VoltOpt  = flag.Bool("v", false, "help message for \"v\" option")
	HalfOpt  = flag.Bool("h", false, "help message for \"h\" option")
	XOpt  = flag.Bool("x", false, "help message for \"x\" option")
)

func main() {

	Settings.ReadSettingsFile()

	flag.Parse()
	if *ChangeOpt {
		RecentMailDateTime := Receiver.GetRecentMailDateTime(Settings.SettingsXml.Config.MailtextPath)
		isAction, AdjustmentSec := Analyzer.GetSettingsSec(RecentMailDateTime,Settings.SettingsXml.Config)
		if isAction {
			log.Printf("Time adjustment was executed. (sec:%d)",  AdjustmentSec)
			Sender.SendSettingMin(AdjustmentSec, Settings.SettingsXml.Smtp)
		}else{
			log.Println("Time adjustment could not be performed.")
		}
	} else if *LateOpt{

	}else if *MailAddressOpt{

	}else if *IdOpt{

	}else if *Offset1Opt{

	}else if *Offset2Opt{

	}else if *Offset3Opt{

	}else if *Offset4Opt{

	}else if *VoltOpt{

	}else if *HalfOpt{

	}else if *XOpt{

	}

}
