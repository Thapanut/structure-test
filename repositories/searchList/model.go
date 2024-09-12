package repositories

import (
	"time"

	"github.com/google/uuid"
)

type SearchListDao struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	DocId           string    `json:"doc_id" gorm:"type:varchar(50)"`
	KeyName         string    `json:"key_name" gorm:"type:varchar(1000)"`
	BirthDate       string    `json:"birth_date" gorm:"type:varchar(10)"`
	NationalityCode *string   `json:"nationality_code" gorm:"type:varchar(2)"`
	FirstName       string    `json:"first_name" gorm:"type:varchar(1000)"`
	LastName        string    `json:"last_name" gorm:"type:varchar(300)"`
	DataSource      string    `json:"data_source" gorm:"type:varchar(30)"`
	CreatedDate     time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate     time.Time `json:"updated_date" gorm:"autoUpdateTime"`
	CreatedBy       string    `json:"created_by" gorm:"type:varchar(100)"`
	CreatedByName   string    `json:"created_by_name" gorm:"type:varchar(100)"`
	UpdatedBy       string    `json:"updated_by" gorm:"type:varchar(100)"`
	UpdatedByName   string    `json:"updated_by_name" gorm:"type:varchar(100)"`
	IsActive        string    `json:"is_active" gorm:"default:Y;type:varchar(1)"`
}

func (k *SearchListDao) TableName() string {
	return "search_list"
}
