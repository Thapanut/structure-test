package repositories

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type IServiceSearchListDao interface {
	GetSearchListByDocIdDataSourceDao(docId string) (resp *SearchListDao, err error)
}

type serviceSearchListDao struct {
	db *gorm.DB
}

func NewSearchListDaoServices(db *gorm.DB) IServiceSearchListDao {
	return &serviceSearchListDao{db: db}
}

func (r *serviceSearchListDao) GetSearchListByDocIdDataSourceDao(docId string) (*SearchListDao, error) {
	var searchList *SearchListDao
	if err := r.db.Where("doc_id = ? and is_active = 'Y'", docId).First(&searchList).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			log.Error("Not Found Search List\n")
			return nil, nil
		}
		log.Errorf("Query Search List By Doc Id error %v\n", err)
		return nil, err
	}
	return searchList, nil
}
