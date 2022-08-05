package models

import (
	"fmt"
	"math"
)

type PaginationResp struct {
	Size      i        nt `json:"size`
	Page      i        nt `json:"page`
	TotalPage i        nt `json:"page_counts"`
	PageIndex i        nt `json:"page_index"`
	D     ata interface{} `json:"data_list`
}

type Pagination struct {
	Size int    `json:"size"`
	Page int    `json:"page"`
	Sort string `json:"sort"`
}

func (q *Pagination) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

// // db querry offet, page index
// func GetOffSet(pageIndex int, limit int, pages int) int {
// 	if pageIndex > pages {
// 		pageIndex = pages
// 	}
// 	offSet := (pageIndex - 1) * limit
// 	return offSet
// }

// Get limit
func (q *Pagination) GetLimit() int {
	return q.Size
}

// // db querry limit, page size
// func GetLimit(size int) int {
// 	if size == 0 {
// 		size = DefaultLimit
// 	}
// 	return size
// }

// Get OrderBy
func (q *Pagination) GetSort() string {
	return q.Sort
}

func (q *Pagination) GetPage() int {
	return q.Page
}

// Get OrderBy
func (q *Pagination) GetSize() int {
	return q.Size
}

func (q *Pagination) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&sort=%s", q.GetPage(), q.GetSize(), q.GetSort())
}



// Get total pages int
func GetTotalPages(totalCount int64, size int) int {
	if size == 0 {
		size = DefaultLimit
	}

	pages := totalCount / size

	if totalCount%pageSize > 0 {
		pages++
	}
	return pages
}

// Get has more
func GetHasMore(currentPage int, totalCount int, pageSize int) bool {
	return currentPage < totalCount/pageSize
}



package models

import "fmt"

const (
	DefaultLimit int = 25
	// DefaultOffset int = 0
)

type Core struct {
	Size       int `json:"size"`
	PageIndex  int `json:"page_index"`
	PageCounts int `json:"page_counts"`
	Total      int `json:"total"`
}

type PaginationInterface interface {
	GetPages()
	GetOffSet()
	GetLimit()
	GetCurrentPage()
	SetData()
}
type Pagination struct {
	Core
	Data interface{} `json:"data"`
}

func NewPaination() *Pagination {
	p := new(Pagination)
	p.PageCounts = 10
	p.PageIndex = 0
	return p
}





