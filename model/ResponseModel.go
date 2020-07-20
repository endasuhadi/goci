package model

type ResponseDataHeader struct {
	ResponseData ResponseData
}

type ResponseData struct {
	ResponseCode    int `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}