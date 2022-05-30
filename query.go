package main

import (
	"fmt"
	"net/url"
	"strconv"
)

type Query struct {
	Country  string
	Category string
	Sources  string
	Keywords string
	Language string
	PageSize int
}

func (q *Query) Values() (url.Values, error) {
	values := url.Values{}

	if q.Sources != "" && (q.Category != "" || q.Country != "") {
		return values, fmt.Errorf("cannot mix sources with the country  or category")
	}

	if q.Country != "" {
		values.Add("country", q.Country)
	}
	if q.Category != "" {
		values.Add("category", q.Category)
	}
	if q.Sources != "" {
		values.Add("sources", q.Sources)
	}
	if q.Keywords != "" {
		values.Add("q", q.Keywords)
	}
	if q.Language != "" {
		values.Add("language", q.Language)
	}
	if q.PageSize != 20 {
		values.Add("pageSize", strconv.Itoa(q.PageSize))
	}

	return values, nil
}

func (q *Query) Validate() error {
	if q.PageSize < 1 && q.PageSize > 100 {
		return fmt.Errorf("number of articles is out of range")
	}
	return nil
}
