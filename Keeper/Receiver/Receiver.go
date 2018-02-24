// Copyright (c) 2018 SHIGEMUNE Katsuhiro
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
package  Receiver

//Module:300
import (
	"time"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
	_"fmt"
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

func setDateTime(dt string) (time.Time,bool){
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dummy:= time.Date(1900, 1, 1, 0, 0, 0, 0, loc)

	d1,d2,d3,t1,t2 :=0,0,0,0,0
	//Date: Fri, 2 Sep 2016 16:36:54 +0900
	c1:=strings.Split(dt," ")
	if len(c1) < 6 {
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
}
	d1, _ = strconv.Atoi(c1[4])	//2016
	d2 = Month2Number(c1[3])	//Sep
	d3, _ = strconv.Atoi(c1[2])	//2

	if d1>=1970 && d1<=2200 {
	}else{
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
	}
	if d2>=1 && d2<=12 {
	}else{
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
	}
	if d3>=1 && d3<=31 {
	}else{
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
	}

	t:=strings.Split(c1[5],":")	//16:36:54
	if len(t) < 2 {
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
	}

	t1, _ = strconv.Atoi(t[0])	//16
	t2, _ = strconv.Atoi(t[1])	//36
	if t1>=0 && t1<=24 {
	}else{
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
	}
	if t2>=0 && t2<=59 {
	}else{
		log.Printf(",304,受信日時に採用したデータに異常があります。%d/%d/%d %d:%d\n",d1,d2,d3,t1,t2)
		return dummy,true
	}

	return time.Date(d1, time.Month(d2), d3, t1, t2, 0, 0, loc),false
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func newFile(fn string) *os.File {
	fp, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	if err != nil {
		log.Println(",300,受信データファイルの展開中にエラーが発生しました。")
	}
	return fp
}

func GetRecentMailDateTime(path string) (time.Time,time.Time,bool){

	loc, _ := time.LoadLocation("Asia/Tokyo")
	dummy := time.Date(1900, 1, 1, 0, 0, 0, 0, loc)
	fp := newFile(path)
	defer fp.Close()
	reader := bufio.NewReader(fp)
	str, err := Readln(reader)
	SendingDateNew :=""
	SendingDateOld :=""
	if err != nil {
		log.Println(",301,受信データファイルの読込中にエラーが発生しました。")
	}
	for err == nil {
		//fmt.Println(str)
		str, err = Readln(reader)
		if 5 <= len(str) {
			if "Date:" == str[0:5] {
				//Date: Thu, 1 Feb 2018 18:58:57 +0900
				SendingDateOld = SendingDateNew
				SendingDateNew = str
			}
		}
	}

	log.Printf(",302,最後に受信したメールの送信日付は、%s でした。",SendingDateNew )

	if ""==SendingDateNew || "" ==SendingDateOld{
		log.Println(",303,読み込んだファイルに日時データがありませんでした。" )
		return dummy,dummy,true
	}
	new,hasErr:=setDateTime(SendingDateNew)
	old,hasErr:=setDateTime(SendingDateOld)
	if hasErr{
		return dummy,dummy,true
	}
	return new,old,false
}