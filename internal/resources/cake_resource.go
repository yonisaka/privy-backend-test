package resources

import (
	"encoding/json"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/helpers"
)

type cakeResource struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func CakeResource(data interface{}) interface{} {
	switch result := data.(type) {
	default:
		panic("Param CakeResource undefined")
	case *[]domain.Cake:
		resultProcessed := MultiCakeResource(result)

		response := struct {
			Data   interface{} `json:"data"`
			Status int         `json:"status"`
			Error  int         `json:"error"`
		}{
			Data:   resultProcessed,
			Status: 200,
			Error:  0,
		}

		return response
	case *domain.Cake:
		resultProcessed := SingleCakeResource(result)

		response := struct {
			Data   interface{} `json:"data"`
			Status int         `json:"status"`
			Error  int         `json:"error"`
		}{
			Data:   resultProcessed,
			Status: 200,
			Error:  0,
		}

		return response
	}

}

func MultiCakeResource(data interface{}) []interface{} {

	result := []map[string]interface{}{}
	jsonByte, _ := json.Marshal(data)
	json.Unmarshal(jsonByte, &result)

	arrData := make([]interface{}, 0)

	for _, value := range result {
		arrData = append(arrData, SingleCakeResource(value))
	}

	return arrData
}

func SingleCakeResource(data interface{}) interface{} {
	dataResult := cakeResource{}
	jsonByte, _ := json.Marshal(data)
	json.Unmarshal(jsonByte, &dataResult)

	dataResult.Image = helpers.ConvertToImage("cake", dataResult.Image)
	createdAt, _ := helpers.ConvertToTime(dataResult.CreatedAt)
	dataResult.CreatedAt = helpers.ConvertTimeFormat(createdAt, 2).(string)
	updatedAt, _ := helpers.ConvertToTime(dataResult.UpdatedAt)
	dataResult.UpdatedAt = helpers.ConvertTimeFormat(updatedAt, 2).(string)

	return dataResult
}
