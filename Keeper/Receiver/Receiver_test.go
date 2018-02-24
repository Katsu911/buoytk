// Copyright (c) 2018 SHIGEMUNE Katsuhiro
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.
package Receiver

import (
	"testing"
	"time"
)

func TestSetDateTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	actual,actual2 := setDateTime("Date: Fri, 2 Sep 2016 16:36:54 +0900")
	expected, expected2 := time.Date(2016, 9, 2, 16, 36, 0, 0, loc),false
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	dummy := time.Date(1900, 1, 1, 0, 0, 0, 0, loc)
	actual,actual2 = setDateTime("Date:Fri, 2 Sep 2016 16:36:54 +0900")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Frz, 2 Sep 2016 16:36:54 +0900")
	expected, expected2 = time.Date(2016, 9, 2, 16, 36, 0, 0, loc),false
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Fri, 43 Sep 2016 16:36:54 +0900")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Fri, 2 Sev 2016 16:36:54 +0900")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Fri, 2 Sep 12016 16:36:54 +0900")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Fri, 2 Sep 2016 25:36:54 +0900")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Fri, 2 Sep 2016 12:76:54 +0900")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("Date: Fri, 2 Sep 2016 16:36:74 +0900")
	expected, expected2 = time.Date(2016, 9, 2, 16, 36, 0, 0, loc),false
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("2017/3/21 10:21:00")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("2017-3-21 102100")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("'17-3-21T10:21:00")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

	actual,actual2 = setDateTime("2017/3/21T10:21")
	expected, expected2 = dummy,true
	if 0 != actual.Sub(expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	if actual2 != expected2 {
		t.Errorf("got %v\nwant %v", actual2, expected2)
	}

}