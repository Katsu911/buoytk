package  Receiver

import (
	"time"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

func Month2Number(mon string) int{
	month := 0
	switch mon {
	case "Jan": month=1
	case "Feb": month=2
	case "Mar": month=3
	case "Apr": month=4
	case "May": month=5
	case "Jun": month=6
	case "Jul": month=7
	case "Aug": month=8
	case "Sep": month=9
	case "Oct": month=10
	case "Nov": month=11
	case "Dec": month=12
	default: month=0
	}
	return month
}

func setDateTime(dt string) time.Time{

	//Date: Fri, 2 Sep 2016 16:36:54 +0900
	c1:=strings.Split(dt," ")
	//日時
	t:=strings.Split(c1[5],":")

	d1,_ :=strconv.Atoi(c1[4])
	d2:= Month2Number(c1[3])
	d3,_ :=strconv.Atoi(c1[2])
	t1,_ :=strconv.Atoi(t[0])
	t2,_ :=strconv.Atoi(t[1])

	return time.Date(d1, time.Month(d2), d3, t1, t2, 0, 0, time.UTC)
}

func GetRecentMailDateTime(path string) time.Time{

	fp, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	var SendingDate string

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {

		str := scanner.Text()
		if "Date:" == str[0:5] {
			//Date: Thu, 1 Feb 2018 18:58:57 +0900
			SendingDate = scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("RecentSendingDateTime was %s.",SendingDate )

	return setDateTime(SendingDate)
}