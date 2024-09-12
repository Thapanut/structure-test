package searchService

import "github.com/Thapanut/struct-test/models"

type SearchListReq struct {
	SearchByName  string           `json:"name" validate:"omitempty"`
	SearchByDocId string           `json:"doc_id" validate:"omitempty"`
	GroupCode     []string         `json:"group_code" validate:"required"`
	Pageable      *models.Pageable `json:"pageable"`
}
type SearchListRes struct {
	SearchList []SearchRes      `json:"search_list"`
	Pageable   *models.Pageable `json:"pageable"`
}
type SearchSanctionListResponseStandard struct {
	models.ResponseStandard
	Data *SearchListRes `json:"data,omitempty"`
}

type SearchRes struct {
	DocId           string `json:"doc_id"`
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date"`
	NationalityDesc string `json:"nationality_desc"`
	CreatedDate     string `json:"created_date"`
	UpdatedDate     string `json:"updated_date"`
	CreatedBy       string `json:"created_by" `
	UpdatedBy       string `json:"updated_by" `
}
