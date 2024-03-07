package typescript

import (
	"encoding/json"

	"github.com/mikerybka/util"
)

func dedup[T any](arr []T) []T {
	values := []string{}
	for _, item := range arr {
		s := util.JSONString(item)
		dup := false
		for _, v := range values {
			if v == s {
				dup = true
				break
			}
		}
		if !dup {
			values = append(values, s)
		}
	}
	res := []T{}
	for _, v := range values {
		var val T
		json.Unmarshal([]byte(v), &val)
		res = append(res, val)
	}
	return res
}
