package models

type Pageable struct {
	PageSize      int `json:"pageSize" validate:"required"`
	PageNumber    int `json:"pageNumber" validate:"required"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

type ResponseStandard struct {
	Code    string      `json:"code"`
	Status  string      `json:"status"`  // FAIL , SUCCESS , NOT_FOUND
	Message string      `json:"message"` // msg err.Error()
	Data    interface{} `json:"data,omitempty"`
}
