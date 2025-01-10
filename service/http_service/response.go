package http_service

// BaseResponse is the base response struct
type BaseResponse struct {
	ResultCode int64  `json:"result_code" structs:"result_code" mapstructure:"result_code"`
	ResultMsg  string `json:"result_msg,omitempty" structs:"result_msg" mapstructure:"result_msg"`
}

type Response struct {
	BaseResponse
	Data interface{} `json:"data,omitempty" structs:"data" mapstructure:"data"`
	Page interface{} `json:"page,omitempty"`
}
