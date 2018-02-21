package Verifier

import (
	"testing"
	"time"
	"fmt"
)


func TestIsSendingPeriod(t *testing.T){

	for i := 1; i < 7; i++ {
		actual := IsSendingPeriod(i)
		expected := true
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}

	actual := IsSendingPeriod(0)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsSendingPeriod(7)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsNormalMin(t *testing.T) {
	src := "20,21,22,23,24,25,26,27,28,29,50,51,52,53,54,55,56,57,58,59"
	actual := isNormalMin(55, src)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isNormalMin(0, src)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isNormalMin(49, src)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isNormalMin(22, src)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsRecvMailPaster(t *testing.T) {

	loc, _ := time.LoadLocation("Asia/Tokyo")

	tm := time.Now()
	tm2:=tm.Add(-(time.Duration(90)*time.Minute))

	actual := isRecvMailPaster(tm2, 90)
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	tm = time.Date(2018, 1, 4, 16, 42, 0, 0, loc)

	actual = isRecvMailPaster(tm, 90)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}

func TestIsAlreadySendSettingMail(t *testing.T) {

	tm := time.Now()
	tm2:=tm.Add(-(time.Duration(30)*time.Minute))
	str := fmt.Sprintf("%d,%d,%d,%d,%d,%d", tm2.Year(),tm2.Month(),tm2.Day(),tm2.Hour(),tm2.Minute(),tm2.Second())

	actual := isAlreadySendSettingMail(str, true)
	expected :=  true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	tm2 =tm.Add(-(time.Duration(120)*time.Minute))
	str = fmt.Sprintf("%d,%d,%d,%d,%d,%d", tm2.Year(),tm2.Month(),tm2.Day(),tm2.Hour(),tm2.Minute(),tm2.Second())

	actual = isAlreadySendSettingMail(str, true)
	expected =  false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = isAlreadySendSettingMail("", true)
	expected =  true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}

func TestIsLateValue(t *testing.T) {

	actual := IsLateValue(-1800)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsLateValue(1800)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsLateValue(1801)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsLateValue(-1801)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestisNormalSendingInterval(t *testing.T) {

	loc, _ := time.LoadLocation("Asia/Tokyo")
	old := time.Date(2018, 2, 21, 12, 00, 0, 0, loc)
	new := time.Date(2018, 2, 21, 12, 30, 0, 0, loc)

	actual:= isNormalSendingInterval(old, new,25)
	expected := true
	if  actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	old = time.Date(2018, 2, 21, 12, 00, 0, 0, loc)
	new = time.Date(2018, 2, 21, 12, 15, 0, 0, loc)

	actual= isNormalSendingInterval(old, new,25)
	expected = false
	if  actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

}

func TestIsLateDateTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	old := time.Date(2017, 12, 31, 12, 23, 0, 0, loc)
	new := time.Date(2017, 12, 31, 12, 53, 0, 0, loc)
	dummy := time.Date(1900, 1, 1, 0, 0, 0, 0, loc)
	actual,actual2, actual3 := IsLateDateTime("2017-12-31T12:23:21,2017-12-31T12:53:21")
	expected, expected2, expected3 := old,new,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if 0 != actual2.Sub(expected2) {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}
	if actual3 != expected3 {
		t.Errorf("got %v\nwant %v", actual3, expected3)
	}

	actual,actual2,actual3 = IsLateDateTime("2017/3/21 10:21:00,2017/3/21 10:51:00")
	expected, expected2,expected3 = dummy,dummy,false
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if 0 != actual2.Sub(expected2) {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}
	if actual3 != expected3 {
		t.Errorf("got %v\nwant %v", actual3, expected3)
	}

}

func TestIsMailAddress(t *testing.T) {

	actual := IsMailAddress("test@test.com")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsMailAddress("testtest.com")
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	//100文字以上
	actual = IsMailAddress("abcdefghij@abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij.com")
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsID(t *testing.T) {

	actual := IsID("TSUSHIMA BUOY")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	//16文字以上
	actual = IsID("OCEAN PASIFIC PEACE")
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsOffset4Voltage(t *testing.T) {

	actual := IsOffset4Voltage(-1.0)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsOffset4Voltage(1.0)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsOffset4Voltage(1.1)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsOffset4Voltage(-1.1)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsOffset(t *testing.T) {

	actual := IsOffset(-5.0)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsOffset(5.0)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsOffset(5.1)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsOffset(-5.1)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsTerminationVoltage(t *testing.T) {

	actual := IsTerminationVoltage(7.0)
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsTerminationVoltage(12.0)
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsTerminationVoltage(12.1)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = IsTerminationVoltage(6.9)
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
