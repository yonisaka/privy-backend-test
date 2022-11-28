package resources

func GlobalResource(message string) interface{} {
	response := struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
		Error   int    `json:"error"`
	}{
		Message: message,
		Status:  200,
		Error:   0,
	}

	return response
}

func GlobalWithDataResource(data interface{}) interface{} {
	response := struct {
		Data   interface{} `json:"data"`
		Status int         `json:"status"`
		Error  int         `json:"error"`
	}{
		Data:   data,
		Status: 200,
		Error:  0,
	}

	return response
}
