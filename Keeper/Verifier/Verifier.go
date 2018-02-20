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


func IsLateDateTime(dt string)(time.Time, bool){

	d1,d2,d3,t1,t2:=0,0,0,0,0
	loc, _ := time.LoadLocation("Asia/Tokyo")
	// yyyy-MM-ddTHH:mm:ss(ISO8601)
	c1:=strings.Split(dt,"-")
	isNormalTime:=true
	if len(c1) < 3 {isNormalTime=false}
	if isNormalTime{
		d1,_ =strconv.Atoi(c1[0])
		d2,_ =strconv.Atoi(c1[1])
		if d1>=1970 && d1<=2200 {
		}else{isNormalTime=false}
		if d2>=1 && d2<=12 {
		}else{isNormalTime=false}

		c2 := strings.Split(c1[2], "T")
		if len(c2) < 2 {isNormalTime=false}
		if isNormalTime {
			d3, _ = strconv.Atoi(c2[0])
			if d3 >= 1 && d3 <= 31 {
			} else {
				isNormalTime = false
			}

			t := strings.Split(c2[1], ":")
			t1, _ = strconv.Atoi(t[0])
			t2, _ = strconv.Atoi(t[1])
			if t1 >= 0 && t1 <= 24 {
			} else {
				isNormalTime = false
			}
			if t2 >= 0 && t2 <= 59 {
			} else {
				isNormalTime = false
			}
		}
	}

	if isNormalTime{
		return time.Date(d1, time.Month(d2), d3, t1, t2, 0, 0, loc),true
	}
	log.Printf(",206,パイプで受け取った日時データに誤りがあります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
	return time.Date(1900, 1, 1, 0, 0, 0, 0, loc),false
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



func isRecvMailPaster(MailRecvTime time.Time, PeriodMin int)bool{

	duration := time.Since(MailRecvTime)
	MailRecvMinDiff := duration / time.Minute

	if PeriodMin < (int)(MailRecvMinDiff) {
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
	if "" == tmstr {return true} //送信履歴ファイルなし
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
			log.Printf("MailRecvMinDiff:%d\n", MailRecvMinDiff)
			if MailRecvMinDiff > MAIL_SENDING_PERIOD_MIN {
				log.Printf(",205,設定メールを送信済みです。%d 分以上間隔をあけて実行してください。\n", MAIL_SENDING_PERIOD_MIN)
				return false
			}
		}
	}
	return true
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

func GetSettingsSec(t time.Time, d Settings.Config)(bool, int){

	//最新の受信メールの送信時間は設定変更すべき時間帯(分)のものであるか？
	isAction := true
	if isNormalMin(t.Minute(), d.AllowanceMinList){isAction=false}

	//最新のメール受信データが想定よりも過去のメールか？
	if isRecvMailPaster(t, d.TrialPeriod){isAction=false}

	//既に設定メールを送っていないか？
	//2018,2,4,15,0,0
	tmstr,isOK:=getRecentFile("SendingHistoryOfSettingMail")
	if isAlreadySendSettingMail(tmstr,isOK){isAction=false}

	return isAction,getAdjustmentMin(t.Minute(),d.SendMin)*60
}