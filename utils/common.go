package utils

import (
	"math/rand"
	"regexp"
	"runtime"
	"time"
)

const (
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	YYYY_MM_DD_00_00_00 = "2006-01-02 00:00:00"
	YYYY_MM_DD_HH_MM    = "2006年01月02日 15:04"
	ORIGIN_DATE         = "2006-01-02 15:04:05"
	YYYY_MM_DD          = "2006-01-02"
	YYYY_MM_DD_CH       = "2006年01月02日"
	YYYY__MM__DD        = "2006_01_02"
	TIME_HOUR           = 3600
	MOBILE_REGULAR      = `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
)

func GetFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		return f.Name()
	}
	return ""
}

func DateStrToTime(dateStr, formate string) (sec int64) {
	t, _ := time.ParseInLocation(formate, dateStr, time.Local)
	sec = t.Unix()
	return
}

func GetAfterDate(sec int) (timeFmt time.Time, err error) {
	str := GetAfterDateStr(sec, ORIGIN_DATE)
	timeFmt, err = time.ParseInLocation(ORIGIN_DATE, str, time.Local)
	return
}

func StrToDate(dateStr string) (timeFmt time.Time, err error) {
	timeFmt, err = time.ParseInLocation(ORIGIN_DATE, dateStr, time.Local)
	return
}

func FormatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

func Time() int64 {
	return time.Now().Unix()
}

func GenerateRandnum(min, max int) (rnd int) {
	if min >= max || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//验证手机号码
func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(MOBILE_REGULAR)
	return reg.MatchString(mobileNum)
}
