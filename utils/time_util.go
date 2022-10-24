package utils

import (
	"fmt"
	"strconv"
	"time"
)

var cstZone = time.FixedZone("CST", 8*3600)

var DateTimeFormatMap = map[string]string{
	"Y": "2006",
	"y": "06",
	"m": "01",
	"d": "02",
	"H": "15",
	"i": "04",
	"s": "05",
}

// GetTimeLeftToday 今日剩余秒数
func GetTimeLeftToday() int {
	timeStr := time.Now().In(cstZone).Format(YYYY_MM_DD_CH)
	t, _ := time.Parse(YYYY_MM_DD_CH, timeStr)
	timeNumber := t.Unix()
	return int(time.Now().In(cstZone).Unix() - timeNumber)
}

// GetTimeInt get not time str
func GetTimeInt() int64 {
	return time.Now().In(cstZone).Unix()
}

func GetAfterDateStr(sec int, format string) string {
	t := time.Now().Unix()
	t = t - int64(sec)
	return time.Unix(t, 0).Format(format)
}

// GetDateFormat 获取当前格式化时间
func GetDateFormat(dateStr string) string {
	data := ""
	for _, v := range dateStr { // i 是字符的字节位置，v 是字符的拷贝
		s := fmt.Sprintf("%c", v)
		if item, ok := DateTimeFormatMap[s]; ok {
			data = data + item
		} else {
			data = data + s
		}
	}
	return time.Now().In(cstZone).Format(data)
}

// GetDateIntFormat 获取当前格式化时间int型
func GetDateIntFormat(dateStr string) int64 {
	data := ""
	for _, v := range dateStr { // i 是字符的字节位置，v 是字符的拷贝
		s := fmt.Sprintf("%c", v)
		if item, ok := DateTimeFormatMap[s]; ok {
			data = data + item
		}
	}
	t, err := strconv.ParseInt(time.Now().In(cstZone).Format(data), 10, 64)
	if err != nil {
		return 0
	}
	return int64(t)
}

// DateToTime date转时间戳
func DateToTime(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, _ := time.ParseInLocation(ORIGIN_DATE, date, loc)
	return tt.Unix()
}
