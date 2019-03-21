package wvparser

import (
	"strconv"
	"strings"
)

type (
	Element struct {
		Name   string
		Values map[float64]float64
	}
	TextElement string
)


// ElementParse は文字列から名前とKey-Value-Pairを取り出しElementで返す
// returns: Element, error
func (te TextElement) ElementParse() (Element, error) {
	// TextElementのFormat
	// # [Name]
	// key1 , value1
	// key2 , value2
	// ...

	lines := strings.Split(string(te), "\n")
	var rt Element
	rt.Values=make(map[float64]float64)
	for i,v := range lines {
		if i==0 {
			rt.Name=strings.Trim(v,"# ")
		} else {
			kv := strings.Split(strings.ReplaceAll(v," ",""), ",")
			key,err := strconv.ParseFloat(kv[0],64)
			if err != nil {
				return rt, err
			}

			value,err := strconv.ParseFloat(kv[1],64)
			if err != nil {
				return rt, err
			}

			rt.Values[key]=value
		}
	}

	return rt, nil
}
