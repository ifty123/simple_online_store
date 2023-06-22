package dto

import (
	"math"
)

type Pagination struct {
	Page     *int `query:"page" json:"page" example:"1"`
	PageSize *int `query:"page_size" json:"page_size" example:"10"`
}

type SearchGetRequest struct {
	Pagination
	Search   string `query:"search"`
	Category uint   `query:"category"`
}

type SearchGetResponse[T any] struct {
	Data           []T            `json:"data"`
	PaginationInfo PaginationInfo `json:"pagination_info"`
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count" example:"20"`
	MoreRecords bool `json:"more_records,omitempty"`
	TotalPage   int  `json:"total_page" example:"2"`
}

type ByIDRequest struct {
	ID uint `param:"id" validate:"required"`
}

func GetLimitOffset(p *Pagination) (limit, offset int) {

	if p.PageSize != nil {
		limit = *p.PageSize
	} else {
		limit = 10
		p.PageSize = &limit
	}

	if p.Page != nil {
		offset = (*p.Page - 1) * limit
	} else {
		offset = 0
	}

	return
}

func CheckInfoPagination(p *Pagination, count int64) *PaginationInfo {
	info := PaginationInfo{
		Pagination: p,
	}
	var page int
	if p.Page != nil {
		page = *p.Page
	} else {
		page = 1
	}
	info.Page = &page

	info.Count = int(count)
	info.TotalPage = int(math.Ceil(float64(count) / float64(*p.PageSize)))
	info.MoreRecords = true
	if *p.Page >= info.TotalPage {
		info.MoreRecords = false
	}

	return &info
}
