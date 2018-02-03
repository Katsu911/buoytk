package Analyzer

import( "time"
"../Settings"
	"strings"
	"strconv"
	"os"
	"log"
	"bufio"
)

const MAIL_SENDING_PERIOD_MIN = 90

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

func isAlreadySendSettingMail()bool{

	now := time.Now()
	var MailRecvMinDiff int

	file, err := os.Open(`SendingHistoryOfSettingMail`)
	if err != nil {
		log.Println("An error occurred in the endingHistoryOfSettingMail.")
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			log.Println("An error occurred in the endingHistoryOfSettingMail.")
			break
		}
		RecentRecvMail := strings.Split(sc.Text(), ",")
		y,_ :=strconv.Atoi(RecentRecvMail[0])
		m,_ :=strconv.Atoi(RecentRecvMail[1])
		mm := time.Month(m)
		d,_ :=strconv.Atoi(RecentRecvMail[2])
		h,_ :=strconv.Atoi(RecentRecvMail[3])
		mmm,_ :=strconv.Atoi(RecentRecvMail[4])
		s,_ :=strconv.Atoi(RecentRecvMail[5])

		day := time.Date(y,mm, d, h, mmm, s, 0, time.Local)
		duration := now.Sub(day)

		hours0 := int(duration.Hours())
		days := hours0 / 24
		hours := hours0 % 24
		mins := int(duration.Minutes()) % 60

		MailRecvMinDiff = (days*24*60)+(hours*60)+mins
	}

	if MailRecvMinDiff > MAIL_SENDING_PERIOD_MIN {
		return false
	}
	return true
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