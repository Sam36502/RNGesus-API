package dto

type FloatResponse struct {
	Number float64 `json:"num"`
}

type IntResponse struct {
	Number int64 `json:"num"`
}

type MessageResponse struct {
	Message string `json:"msg"`
}
