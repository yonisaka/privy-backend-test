package helpers

import (
	"log"
	"reflect"
	"time"
)

/*
Format Type
1 = "2006-01-02 15:04:05",
2 = "2006-01-02 15:04",
3 = "02 Jan 2006 15:04",
4 = "2006-01-02",
5 = "02 Jan 2006",
6 = "02/01/2006"
7 = "15:04:05",
8 = "15:04"
9 = "2006-01-02T15:04:05"
*/
func ConvertTimeFormat(dateTime interface{}, formatType int) interface{} {
	if dateTime == nil {
		return ""
	}

	var err error
	dateTime, err = ConvertToTime(dateTime)
	if err != nil {
		log.Println("Param  is not time.Time; Param Type : " + reflect.TypeOf(dateTime).String())
		return ""
	} else if dateTime.(time.Time).Format("2006-01-02 15:04") == "0001-01-01 00:00" {
		return ""
	}

	data := dateTime.(time.Time)
	switch formatType {
	default:
		panic("Param formatType is not type")
	case 1:
		return data.Format("2006-01-02 15:04:05")
	case 2:
		return data.Format("2006-01-02 15:04")
	case 3:
		return data.Format("02 Jan 2006 15:04")
	case 4:
		return data.Format("2006-01-02")
	case 5:
		return data.Format("02 Jan 2006")
	case 6:
		return data.Format("02/01/2006")
	case 7:
		return data.Format("15:04:05")
	case 8:
		return data.Format("15:04")
	case 9:
		return data.Format("2006-01-02T15:04:05")
	}
}
