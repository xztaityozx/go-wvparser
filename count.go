package wvparser

import (
	"golang.org/x/xerrors"
	"gonum.org/v1/gonum/unit"
	"log"
	"strconv"
	"strings"
)

type Counter struct {
	Filters map[float64]string
}

func NewCounter(f ...string) Counter {
	rt := Counter{
		Filters: map[float64]string{},
	}

	for _, v := range f {
		err := rt.AddFilter(v)
		if err != nil {
			log.Fatal(xerrors.Errorf("Failed NewCounter: %w", err))
		}
	}

	return rt
}

// AddFilter は数え上げのFilterを追加します
// in: filter string
// 		format: [key[SI unit [G,M,k,m,u,n,p]]:[operator <=,<,>=,>,==,!=]:[value[SI unit [G,M,k,m,u,n,p]]
// returns: error
func (c *Counter) AddFilter(q string) error {
	key, st, err := ParseFilter(q)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	c.Filters[key] = st
	return nil
}

// ParseFilter は数え上げのFilterを文字列から解析する
// in: filter string
// returns: key, statement, error
func ParseFilter(q string) (key float64, statement string, err error) {
	box := strings.Split(q, ":")

	key = ToFloat64(box[0])
	if err != nil {
		return 0, "", xerrors.Errorf(": %w", err)
	}
	statement = box[1] + ":" + box[2]
	return
}

// ToFloat64 はSI単位系をFloat64で返します
func ToFloat64(value string) float64 {

	u := value[len(value)-1:]
	var e float64
	if u == "G" {
		e = unit.Giga
	} else if u == "M" {
		e = unit.Mega
	} else if u == "k" {
		e = unit.Kilo
	} else if u == "m" {
		e = unit.Milli
	} else if u == "u" {
		e = unit.Micro
	} else if u == "n" {
		e = unit.Nano
	} else if u == "p" {
		e = unit.Pico
	} else {
		e = 1
	}

	if e != 1 {
		value = value[:len(value)-1]
	}
	a, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(xerrors.Errorf("failed ToFloat64: %w", err))
	}

	return a * e
}

// Aggregate はCounter structを使ってデータの数え上げを行います
func (c Counter) Aggregate(csv WVCsv) int64 {
	var rt int64 = 0

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
		if status {
			rt++
		}
	}

	return rt
}
