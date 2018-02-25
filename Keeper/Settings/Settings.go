// Copyright (c) 2018 SHIGEMUNE Katsuhiro
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
package Settings

import (
	"io/ioutil"
	"encoding/xml"
)

type Settings struct {
Smtp Smtp `xml:"Smtp"`
Config Config `xml:"Config"`
}

type Smtp struct {
name string `xml:"name"`
Account string `xml:"account"`
Password string `xml:"password"`
Host string `xml:"host"`
Address string `xml:"address"`
Source string `xml:"source"`
}

type Config struct {
SendMin          int    `xml:"send_min"`
AllowanceMinList string `xml:"allowance_min_list"`
RecentRecvPeriod int    `xml:"recent_recv_period"`
TxInterval       int    `xml:"tx_interval"`
MailtextPath     string `xml:"mailtext_path"`
LogFolderPath	 string `xml:"logfolder_path"`
}

var SettingsXml Settings

func ReadSettingsFile() bool {
	SettingsXml = Settings{Smtp{"dummy","","","","",""},
		Config{0,"",0,0,"",""}}
	data, _ := ioutil.ReadFile(`./settings.xml`)
	err := xml.Unmarshal([]byte(data), &SettingsXml)
	if err != nil {return false}
	return true
}