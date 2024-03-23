package helpers

import (
	"fmt"
	"net/url"
	"strings"
)

type Ordering struct {
	Key       string
	Direction string
}

func (o Ordering) String() string {
	return fmt.Sprint(o.Key, " ", o.Direction)
}

func ExtractOrdering(baseURL *url.URL) []Ordering {
	ordering := make([]Ordering, 0)

	for _, field := range strings.Split(baseURL.Query().Get("ordering"), ",") {
		if field != "" {
			key, found := strings.CutPrefix(field, "-")

			if found {
				ordering = append(ordering, Ordering{Key: key, Direction: "DESC"})
			} else {
				ordering = append(ordering, Ordering{Key: field, Direction: "ASC"})
			}
		}
	}

	return ordering
}
