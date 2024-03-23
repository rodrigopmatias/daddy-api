package helpers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func ExtractPaginateInfo(baseURL *url.URL) (int, int) {
	var limit int
	var offset int

	value := baseURL.Query().Get("offset")
	if value == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(value)
	}

	value = baseURL.Query().Get("limit")
	if value == "" {
		limit = 30
	} else {
		limit, _ = strconv.Atoi(value)
	}

	return offset, limit
}

func Paginate(baseURL *url.URL, count int64, offset int, limit int) (*string, *string) {
	return paginateNext(baseURL, count, offset, limit), paginatePrevius(baseURL, offset, limit)
}

func paginatePrevius(baseURL *url.URL, offset int, limit int) *string {
	if offset > 0 {
		var newOffset int64

		newOffset = int64(offset - limit)
		if newOffset < 0 {
			newOffset = 0
		}

		params := make([]string, 0)
		for attr := range baseURL.Query() {
			if attr != "offset" {
				params = append(params, fmt.Sprintf("%s=%s", attr, url.QueryEscape(baseURL.Query().Get(attr))))
			}
		}

		params = append(params, fmt.Sprintf("offset=%d", newOffset))
		newURL := fmt.Sprintf("%s?%s", strings.Split(baseURL.RequestURI(), "?")[0], strings.Join(params, "&"))

		return &newURL
	}

	return nil
}

func paginateNext(baseURL *url.URL, count int64, offset int, limit int) *string {
	if count > int64(offset+limit) {
		var newOffset int64

		newOffset = int64(offset + limit)
		if newOffset > count {
			newOffset = count
		}

		params := make([]string, 0)
		for attr := range baseURL.Query() {
			if attr != "offset" {
				params = append(params, fmt.Sprintf("%s=%s", attr, url.QueryEscape(baseURL.Query().Get(attr))))
			}
		}

		params = append(params, fmt.Sprintf("offset=%d", newOffset))
		newURL := fmt.Sprintf("%s?%s", strings.Split(baseURL.RequestURI(), "?")[0], strings.Join(params, "&"))

		return &newURL
	}

	return nil
}
