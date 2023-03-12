package dto

type Response struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	ErrorMsg string      `json:"errorMsg"`
	Total    int64       `json:"total"`
}

func Ok() Response {
	return Response{
		Success: true,
	}
}

func OkData(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func Err(errStr string) Response {
	return Response{
		Success:  false,
		ErrorMsg: errStr,
	}
}
