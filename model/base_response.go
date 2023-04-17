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

type HomeInformationResponse struct {
	About    string `json:"About"`
	Name     string `json:"Name"`
	Github   string `json:"Github"`
	LinkedIn string `json:"LinkedIn"`
}
