package jsonmodel

// ResponseBase struct
type ResponseBase struct {
	HasError bool        `json:"hasError"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}
