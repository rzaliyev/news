package main

import (
	"fmt"
	"net/url"
)

type Query struct {
	Country  string
	Category string
	Sources  string
	Keywords string
	Language string
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

	return values, nil
}
