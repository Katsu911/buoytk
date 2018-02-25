// Copyright (c) 2018 SHIGEMUNE Katsuhiro
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
package Verifier
//Module:200
import( "time"
"../Settings"
	"strings"
	"strconv"
	"os"
	"log"
	"bufio"
	"regexp"
)


func isSendingIntervalAbnormal( old time.Time, new time.Time, ival int)bool{

	duration := new.Sub(old)
	MailRecvMinInterval := duration / time.Minute

	if ival < (int)(MailRecvMinInterval) {
		log.Println(",208,最新の受信メールと1つ前の受信メールの送信間隔が短すぎます。(=想定外の時間帯にブイの電源が入ってメール送信した可能性がある。)")
		return true
	}
	return false

}

func IsLateDateTime(dt string)(time.Time, time.Time, bool){
	parts := make([][]int, 2)
	for i := range parts {
		parts[i] = make([]int, 5)
	}
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dummyTime := time.Date(1900, 1, 1, 0, 0, 0, 0, loc)
	// yyyy-MM-ddTHH:mm:ss,yyyy-MM-ddTHH:mm:ss		//※フォーマットは、ISO8601形式
	// (1)最新から2番目に古い受信メール、(2) 最新の受信メール
	tmp:=strings.Split(dt,",")	//yyyy-MM-ddTHH:mm:ss,yyyy-MM-ddTHH:mm:ss
	if len(tmp) < 2{
		log.Println(",207,パイプで受け取った日時データに誤りがあります。")
		return dummyTime,dummyTime,false
	}
	for i:=0;i<2;i++ {
		c1 := strings.Split(tmp[i], "-")	//yyyy-MM-ddTHH:mm:ss
		if len(c1) < 3 {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}
		parts[i][0], _ = strconv.Atoi(c1[0])
		parts[i][1], _ = strconv.Atoi(c1[1])
		if parts[i][0] >= 1970 && parts[i][0] <= 2200 {
		} else {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}
		if parts[i][1] >= 1 && parts[i][1]  <= 12 {
		} else {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}

		c2 := strings.Split(c1[2], "T")	//ddTHH:mm:ss
		if len(c2) < 2 {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}
		parts[i][2], _ = strconv.Atoi(c2[0])	//dd
		if parts[i][2] >= 1 && parts[i][2] <= 31 {
		} else {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}

		t := strings.Split(c2[1], ":")	//HH:mm:ss
		parts[i][3], _ = strconv.Atoi(t[0])
		parts[i][4], _ = strconv.Atoi(t[1])
		if parts[i][3] >= 0 && parts[i][3] <= 24 {
		} else {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}
		if parts[i][4] >= 0 && parts[i][4] <= 59 {
		} else {
			log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",parts[i][0],parts[i][1],parts[i][2],parts[i][3],parts[i][4])
			return dummyTime,dummyTime,false
		}
	}

	old := time.Date(parts[0][0], time.Month(parts[0][1]), parts[0][2], parts[0][3], parts[0][4], 0, 0, loc)
	new := time.Date(parts[1][0], time.Month(parts[1][1]), parts[1][2], parts[1][3], parts[1][4], 0, 0, loc)
	return old,new,true
	}




func isNormalMin(min int, MinList string) bool{

	mins := strings.Split(MinList, ",")
	for _, v := range mins {
		v2, _ := strconv.Atoi(v)
		if min==v2 {
			return true
		}
	}
	log.Printf(",200,検査した(分)は、正常な時間帯(分)のものではありません。(%d分)\n",min)
	return false
}


func IsLateValue(v int)bool{
	const UpperLateRange=1800
	const LowerLateRange=-1800

	if v>=LowerLateRange && v<=UpperLateRange{
		return true
	}
	log.Println(",201,遅延時間が定義域の外にあります。")
	return false
}



func isOperationBuoy(MailRecvTime time.Time, PeriodMin int)bool{

	duration := time.Since(MailRecvTime)
	MailRecvMinDiff := duration / time.Minute

	if PeriodMin >= (int)(MailRecvMinDiff) {
		log.Println(",202,最新のメール受信時間(分)は規定よりも前に送られたものです。(=ブイが稼働していない可能性がある。)")
		return true
	}
	return false
}

func getAdjustmentMin(src int, target int)int{

	if src==target {
		return 0
	}

	min := target-src

	if min>30 {
		min=min-60
	}

	return min
}




func getRecentFile(f string)(string, bool){

	rtn :=""

	file, err := os.OpenFile(f, os.O_RDONLY, 0644)
	if err != nil {
		return "",true
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			log.Println(",203,過去に設定メールが送られたことを示すSendingHistoryOfSettingMailファイルを読み込んだ時にエラーが発生しました。")
			return "",false
			break
		}
		//yyyy,mm,dd,hh,mm,ss
		//2018,2,7,13,47,55
		rtn = sc.Text()
	}
	if "" != rtn {
		return rtn,true
	}
	log.Println(",204,過去に設定メールが送られたことを示すSendingHistoryOfSettingMailファイルの本文が空です。")
	return "1900,1,1,0,0,0",false
}

func isAlreadySendSettingMail(tmstr string, isOK bool)bool {

	const MAIL_SENDING_PERIOD_MIN= 90

	now := time.Now()
	var MailRecvMinDiff int
	if "" == tmstr {return false} //送信履歴ファイルなし
	if isOK {
		RecentRecvMail := strings.Split(tmstr, ",")
		if 6 <= len(RecentRecvMail) {
			y, _ := strconv.Atoi(RecentRecvMail[0])
			m, _ := strconv.Atoi(RecentRecvMail[1])
			mm := time.Month(m)
			d, _ := strconv.Atoi(RecentRecvMail[2])
			h, _ := strconv.Atoi(RecentRecvMail[3])
			mmm, _ := strconv.Atoi(RecentRecvMail[4])
			s, _ := strconv.Atoi(RecentRecvMail[5])
			loc, _ := time.LoadLocation("Asia/Tokyo")
			day := time.Date(y, mm, d, h, mmm, s, 0, loc)
			duration := now.Sub(day)
			MailRecvMinDiff = (int)(duration/time.Minute)
			if MailRecvMinDiff < MAIL_SENDING_PERIOD_MIN {
				log.Printf(",205,設定メールを送信済みです。%d 分以上間隔をあけて実行してください。\n", MAIL_SENDING_PERIOD_MIN)
				return true
			}
		}
	}
	return false
}


func IsSendingPeriod(v int)bool{

	const UpperModeRange=6
	const LowerModeRange=1

	if v>=LowerModeRange && v<=UpperModeRange{
		return true
	}
	return false
}

func IsMailAddress(address string)bool{

	var email_pattern = `^([a-zA-Z0-9])+([a-zA-Z0-9\._-])*@([a-zA-Z0-9_-])+([a-zA-Z0-9\._-]+)+$`
	var email_re = regexp.MustCompile(email_pattern)
	const MAX_MAIL_ADDRESS_LENGTH = 100

	if len(email_re.FindAllString(address, -1)) != 0{
		if len(address) <= MAX_MAIL_ADDRESS_LENGTH {
			return true
		}
	}
	return false
}

func IsID(id string)bool{

	const MAX_ID_LENGTH = 16

	if len(id) <= MAX_ID_LENGTH {
		return true
	}
	return false
}

func IsOffset(v float64)bool{

	const UpperOffsetRange=5.0
	const LowerOffsetRange=-5.0

	if v>=LowerOffsetRange && v<=UpperOffsetRange{
		return true
	}
	return false
}

func IsOffset4Voltage(v float64)bool{

	const UpperOffsetRange=1.0
	const LowerOffsetRange=-1.0

	if v>=LowerOffsetRange && v<=UpperOffsetRange{
		return true
	}
	return false
}

func IsTerminationVoltage(v float64)bool{

	const UpperVoltageRange=12.0
	const LowerVoltageRange=7.0

	if v>=LowerVoltageRange && v<=UpperVoltageRange{
		return true
	}
	return false
}

func GetSettingsSec(old time.Time, new time.Time, d Settings.Config)(bool, int){

	//最新の受信メールの送信時間は通常の送信されるべき時間帯(分)のものか？
	//Yes->何もしない。 No->補正する。
	isAction := true
	if isNormalMin(new.Minute(), d.AllowanceMinList){
		log.Println(",208,isNormalMin->true ActionFlag:false")
	isAction=false
	}else{
		log.Println(",209,isNormalMin->false ActionFlag:true")
	}

	//直近の受信メールが古いものでないか？(ブイは稼働しているか？)
	//No->何もしない。
	if !isOperationBuoy(new, d.RecentRecvPeriod){
		log.Println(",210,isOperationBuoy->false ActionFlag:false")
		isAction=false
	}else{
		log.Println(",211,isOperationBuoy->true ActionFlag:true")
	}

	//受信メールと受信メールの送信間隔は異常か？
	//(=試験時などイレギュラな電源の入切によるメール送信ではないか？)
	//Yes->何もしない。
	if isSendingIntervalAbnormal(old,new,d.TxInterval){
		log.Println(",212,isSendingIntervalAbnormal->true ActionFlag:false")
		isAction=false
	}else{
		log.Println(",213,isSendingIntervalAbnormal->false ActionFlag:true")
	}

	//既に設定メールを送っている?
	//Yes->何もしない。
	//2018,2,4,15,0,0
	tmstr,isOK:=getRecentFile("SendingHistoryOfSettingMail")
	if isAlreadySendSettingMail(tmstr,isOK){
		log.Println(",214,isAlreadySendSettingMail->true ActionFlag:false")
		isAction=false
	}else{
		log.Println(",215,isAlreadySendSettingMail->false ActionFlag:true")
	}

	sec:=getAdjustmentMin(new.Minute(),d.SendMin)*60
	log.Printf(",216,isAction=%t, AdjustmentSec=%d \n",isAction,sec)

	return isAction,sec
}