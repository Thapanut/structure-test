package searchService

import (
	"github.com/Thapanut/struct-test/constants/e"
	"github.com/Thapanut/struct-test/errs"
	sl "github.com/Thapanut/struct-test/repositories/searchList"
)

type IServiceSearch interface {
	SearchList(req SearchListReq) (*SearchListRes, error)
}

type SearchService struct {
	searchListRepo sl.IServiceSearchListDao
}

func NewSearchService(
	searchListRepo sl.IServiceSearchListDao,
) IServiceSearch {
	return &SearchService{
		searchListRepo: searchListRepo,
	}
}

// SearchList implements IServiceSearch.
func (s *SearchService) SearchList(req SearchListReq) (*SearchListRes, error) {
	var resp SearchListRes
	searchListDao, err := s.searchListRepo.GetSearchListByDocIdDataSourceDao(req.SearchByDocId)
	if err != nil {
		return nil, errs.AppError{
			Code:    e.INTERNAL_SERVER_ERROR,
			BuCode:  "500",
			Message: err.Error(),
		}
	}
	if searchListDao == nil {
		return nil, errs.AppError{
			Code:    e.SUCCESS,
			BuCode:  "404",
			Message: "data not found",
		}
	}
	res := SearchRes{
		DocId:     searchListDao.DocId,
		BirthDate: searchListDao.BirthDate,
	}
	resp.SearchList = append(resp.SearchList, res)
	return &resp, nil
}
