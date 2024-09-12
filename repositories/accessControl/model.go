package accessControl

type AccessControl struct {
	ID          uint   `json:"id" gorm:"type:uuid;not null"`
	RoleCode    string `json:"role_code" gorm:"type:varchar(50);"`
	APIName     string `json:"api_name" gorm:"type:varchar(100);"`
	HTTPMethod  string `json:"http_method" gorm:"type:varchar(10);"`
	URL         string `json:"url" gorm:"type:varchar(100);"`
	IsActive    string `json:"is_active" gorm:"type:varchar(1);"`
	CreatedDate string `json:"created_date"`
	CreatedBy   string `json:"created_by"`
	UpdatedDate string `json:"updated_date"`
	UpdatedBy   string `json:"updated_by" gorm:"type:varchar(100);"`
}

func (AccessControl) TableName() string {
	return "access_control"
}
