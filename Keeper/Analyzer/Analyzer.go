package Analyzer

import( "time"
"../Settings"
	"strings"
	"strconv"
	"os"
	"log"
	"bufio"
	"regexp"
	"testing"
)


func isRightMin(min int, MinList string) bool{

	mins := strings.Split(MinList, ",")
	for _, v := range mins {
		v2, _ := strconv.Atoi(v)
		if min==v2 {
			return true
		}
	}
	return false
}

func IsLateValue(v int)bool{
	const UpperLateRange=1800
	const LowerLateRange=-1800

	if v>=LowerLateRange && v<=UpperLateRange{
		return true
	}
	return false
}

func TestisRightMin(t *testing.T) {
	src := "50,51,52,53,54,55,56,57,58,59"
	actual := isRightMin(55, src)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isRightMin(0, src)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isRightMin(49, src)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isRightMin(22, src)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func isRecvMailRecent(MailRecvTime time.Time, PeriodMin int)bool{

	now := time.Now()
	duration := now.Sub(MailRecvTime)

	hours0 := int(duration.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60

	MailRecvMinDiff := (days*24*60)+(hours*60)+mins

	if PeriodMin > MailRecvMinDiff {return false}
	return true
}

func TestisRecvMailRecent(t *testing.T) {

	tm := time.Date(2018, 2, 4, 16, 42, 0, 0, time.Local)

	actual := isRecvMailRecent(tm, 90)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	tm = time.Date(2018, 1, 4, 16, 42, 0, 0, time.Local)

	actual = isRecvMailRecent(tm, 90)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}

func getadjustmentMin(src int, target int)int{

	if src==target {
		return 0
	}

	min := target-src

	if min>30 {
		min=min-60
	}
	return min
}


func TestgetadjustmentMin(t *testing.T) {

	actual := getadjustmentMin(49, 52)
	expected := 3
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(0, 52)
	expected = -8
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(30, 52)
	expected = 22
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(8, 52)
	expected = -16
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(49, 58)
	expected = 9
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(0, 58)
	expected = -2
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(30, 58)
	expected = 28
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = getadjustmentMin(8, 58)
	expected = -10
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}

func getRecentFile(f string)(string, bool){

	rtn :=""

	file, err := os.Open(f)
	if err != nil {
		log.Println("An error occurred in the endingHistoryOfSettingMail.")
		return "",false
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			log.Println("An error occurred in the endingHistoryOfSettingMail.")
			return "",false
			break
		}
		rtn = sc.Text()
	}
	if "" != rtn {
		return rtn,true
	}
	return "",false
}

func TestgetRecentFile(t *testing.T) {
	actual1, actual2 := getRecentFile("SendingHistoryOfSettingMail")
	expected1,expected2 := "2018,2,4,15,0,0", true
	if actual1 != expected1 {
		t.Errorf("got %v\nwant %v", actual1, expected1)
	}
	if actual2 != expected2  {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual1, actual2 = getRecentFile("SendingHistoryOfSettingMail_dummy")
	expected1,expected2 = "", false
	if actual1 != expected1 {
		t.Errorf("got %v\nwant %v", actual1, expected1)
	}
	if actual2 != expected2  {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}
}

func isAlreadySendSettingMail()bool {

	const MAIL_SENDING_PERIOD_MIN= 90

	now := time.Now()
	var MailRecvMinDiff int
	//2018,2,4,15,0,0
	tm, isOK := getRecentFile("SendingHistoryOfSettingMail")
	if (isOK) {
		RecentRecvMail := strings.Split(tm, ",")
		y, _ := strconv.Atoi(RecentRecvMail[0])
		m, _ := strconv.Atoi(RecentRecvMail[1])
		mm := time.Month(m)
		d, _ := strconv.Atoi(RecentRecvMail[2])
		h, _ := strconv.Atoi(RecentRecvMail[3])
		mmm, _ := strconv.Atoi(RecentRecvMail[4])
		s, _ := strconv.Atoi(RecentRecvMail[5])

		day := time.Date(y, mm, d, h, mmm, s, 0, time.Local)
		duration := now.Sub(day)

		hours0 := int(duration.Hours())
		days := hours0 / 24
		hours := hours0 % 24
		mins := int(duration.Minutes()) % 60

		MailRecvMinDiff = (days * 24 * 60) + (hours * 60) + mins

		if MailRecvMinDiff > MAIL_SENDING_PERIOD_MIN {
			return false
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
	var isAction = true
	if !(isRightMin(t.Minute(), d.AllowanceMinList)){
		isAction=false
	}
	//最新のメール受信データは最近のものか？
	if !(isRecvMailRecent(t, d.TrialPeriod)){
		isAction=false
	}
	//既に設定メールを送っていないか？
	if !(isAlreadySendSettingMail()){
		isAction=false
	}

	return isAction,getadjustmentMin(t.Minute(),d.SendMin)
}