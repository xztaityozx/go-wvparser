package wvparser

import "strings"

func (c Counter) GetStatuses(csv WVCsv) ([]bool, error) {
	var rt []bool

	f := func(k float64, y float64) bool {
		s := c.Filters[k]
		box := strings.Split(s, ":")

		x := ToFloat64(box[1])

		if box[0] == "<=" {
			return y <= x
		} else if box[0] == "<" {
			return y < x
		} else if box[0] == ">=" {
			return y >= x
		} else if box[0] == ">" {
			return y > x
		} else if box[0] == "==" {
			return y == x
		} else if box[0] == "!=" {
			return y != x
		} else {
			return false
		}
	}

	for _, v := range csv.Data {
		status := true
		for k := range c.Filters {
			status = status && f(k, v.Values[k])
		}
		rt = append(rt, status)
	}

	return rt, nil
}
