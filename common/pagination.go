package common

import (
	"github.com/biezhi/gorm-paginator/pagination"
)

type PaginationFormat struct {
	CurrentPage int         `json:"currentPage"`
	FirstPage   bool        `json:"firstPage"`
	IsNext      bool        `json:"isNext"`
	Items       interface{} `json:"items"`
	LastPage    bool        `json:"lastPage"`
	PageSize    int         `json:"pageSize"`
	TotalItems  int         `json:"totalItems"`
	TotalPage   int         `json:"totalPage"`
}

func ReFormatPagination(s *pagination.Paginator) PaginationFormat {

	return PaginationFormat{
		TotalItems:  s.TotalRecord,
		TotalPage:   s.TotalPage,
		Items:       s.Records,
		CurrentPage: s.Page,
		IsNext:      s.Page != s.TotalPage,
		FirstPage:   s.Page == 1,
		LastPage:    s.Page == s.TotalPage,
		PageSize:    s.Limit,
	}
}
