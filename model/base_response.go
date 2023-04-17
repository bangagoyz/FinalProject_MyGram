package model

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type FailedResponse struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}
