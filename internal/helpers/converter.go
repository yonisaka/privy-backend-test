package helpers

import (
	"errors"
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
