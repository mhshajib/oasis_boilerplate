package utils

import (
	"net/http"
	"strconv"

	"github.com/mhshajib/oasis_boilerplate/pkg/config"
)

// Pagination ...
type Pagination struct {
	CurrentPage int64 `json:"current_page"`
	NextPage    int64 `json:"next_page"`
	Limit       int64 `json:"limit"`
	Count       int64 `json:"count"`
}

// NewPagination ...
func NewPagination(count, page, limit int64) *Pagination {
	if page <= 0 {
		page = 1
	}

	p := &Pagination{
		Count: count,
		Limit: limit,
	}

	p.CurrentPage = page
	p.NextPage = p.CurrentPage + 1
	return p
}

// GetPager return page & limit number from http request
func GetPager(r *http.Request) (page, limit, offset int64, err error) {
	p := r.URL.Query().Get("page")
	l := r.URL.Query().Get("limit")
	maxLimit := int64(config.HttpApp().PaginationLimit)
	limit = maxLimit

	if p != "" {
		page, err = strconv.ParseInt(p, 10, 32)
		if err != nil {
			return
		}
	}
	if l != "" {
		limit, err = strconv.ParseInt(l, 10, 32)
		if err != nil {
			return
		}
		if limit > maxLimit {
			limit = maxLimit
		}
	}
	if page <= 1 {
		page = 1
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	return
}
