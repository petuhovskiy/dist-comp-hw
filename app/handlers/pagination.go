package handlers

import (
	"net/http"
	"strconv"
)

func parseUint(str string) (uint, error) {
	res, err := strconv.ParseUint(str, 10, 64)
	return uint(res), err
}

func parseLimitOffset(r *http.Request) (limit, offset uint, err error) {
	if s := r.URL.Query().Get("limit"); s != "" {
		limit, err = parseUint(s)
		if err != nil {
			return 0, 0, err
		}
	}

	if s := r.URL.Query().Get("offset"); s != "" {
		offset, err = parseUint(s)
		if err != nil {
			return 0, 0, err
		}
	}

	return limit, offset, nil
}
