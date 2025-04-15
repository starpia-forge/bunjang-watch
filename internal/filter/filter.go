package filter

import (
	"github.com/starpia-forge/bunjang-watch/internal/bunjang"
	"regexp"
	"strconv"
)

type Filter[T any] interface {
	Apply(T) bool
}

type StatusFilter struct{}

func (f *StatusFilter) Apply(p bunjang.Product) bool {
	status, err := strconv.Atoi(p.Status)
	if err != nil {
		return false
	}
	return status < bunjang.ProductStatusSoldOut
}

type MinPriceFilter struct {
	MinPrice int
}

func (f *MinPriceFilter) Apply(p bunjang.Product) bool {
	if p.Price == "" {
		return false
	}
	price, err := strconv.Atoi(p.Price)
	if err != nil {
		return false
	}
	return price >= f.MinPrice
}

type MaxPriceFilter struct {
	MaxPrice int
}

func (f *MaxPriceFilter) Apply(p bunjang.Product) bool {
	if p.Price == "" {
		return false
	}
	price, err := strconv.Atoi(p.Price)
	if err != nil {
		return false
	}
	return price <= f.MaxPrice
}

type KeywordFilter struct {
	Keywords []*regexp.Regexp
}

func (f *KeywordFilter) Apply(p bunjang.Product) bool {
	if len(f.Keywords) == 0 {
		return true
	}
	for _, keyword := range f.Keywords {
		if keyword.MatchString(p.Name) {
			return true
		}
	}
	return false
}

type IgnoreKeywordFilter struct {
	IgnoreKeywords []*regexp.Regexp
}

func (f *IgnoreKeywordFilter) Apply(p bunjang.Product) bool {
	for _, keyword := range f.IgnoreKeywords {
		if keyword.MatchString(p.Name) {
			return false
		}
	}
	return true
}

type IncludeUsedFilter struct {
	IncludeUsed bool
}

func (f *IncludeUsedFilter) Apply(p bunjang.Product) bool {
	if f.IncludeUsed {
		return true
	}
	return p.Used == bunjang.ProductUsedNew
}
