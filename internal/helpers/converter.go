package helpers

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ConvertToTime(data interface{}) (time.Time, error) {
	switch dataParse := data.(type) {
	default:
		return time.Time{}, errors.New("invalid data type")
	case string:
		if strings.Contains(dataParse, ".000000") {
			dataParse = strings.ReplaceAll(dataParse, ".000000", "")
		}

		result, err := time.Parse(time.RFC3339, dataParse)
		if err != nil {
			return time.Time{}, err
		}

		return result, nil

	case time.Time:
		return dataParse, nil
	}
}

func ConvertToImage(modul string, filename string) string {
	path := GoDotEnvVariable("BASE_URL") + "/image/" + modul + "/"
	if filename == "" {
		return path + "no-image.jpeg"
	}
	return path + filename
}

func ConvertToString(data interface{}) string {

	if data == nil {
		return ""
	}

	switch data.(type) {
	default:
		panic("Param " + reflect.TypeOf(data).String() + " convert to string undefined")
	case string:
		return data.(string)
	case time.Time:
		return data.(time.Time).Format("2006-01-02 15:04:05")
	case int:
		return strings.TrimRight(strconv.Itoa(data.(int)), "\n")
	case int8:
		return strings.TrimRight(strconv.Itoa(int(data.(int8))), "\n")
	case int32:
		return strings.TrimRight(strconv.Itoa(int(data.(int32))), "\n")
	case uint32:
		return strings.TrimRight(strconv.Itoa(int(data.(uint32))), "\n")
	case int64:
		return strings.TrimRight(strconv.Itoa(int(data.(int64))), "\n")
	case float64:
		return strings.TrimRight(strconv.FormatFloat(data.(float64), 'f', -1, 64), "\n")
	case []uint8:
		return string(data.([]uint8))
	case map[string]interface{}:
		result, _ := json.Marshal(data)
		return string(result)
	case []map[string]interface{}:
		result, _ := json.Marshal(data)
		return string(result)
	case []interface{}:
		result, _ := json.Marshal(data)
		return string(result)
	}
}
