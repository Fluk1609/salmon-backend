package dto

type AllocateRequest struct {
	Orders interface{} `json:"orders"`
}

type AllocateResponse struct {
	Data interface{} `json:"data"`
}