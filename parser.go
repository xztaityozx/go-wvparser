package wvparser

import (
	"golang.org/x/xerrors"
	"io/ioutil"
	"strings"
)

type (
	WVParser struct {
		FilePath string
	}
	WVCsv struct {
		Data   []*Element
		Header Header
	}
	Header struct {
		Format    string
		SavedTime string
		Signal    string
	}
	Document struct {
		Header       Header
		TextElements []TextElement
	}
)

func (p WVParser) Parse() (WVCsv, error) {
	var rt WVCsv

	d, err := p.GetDocument()
	if err != nil {
		return WVCsv{}, xerrors.Errorf(": %w", err)
	}

	rt.Header=d.Header
	rt.Data, err = d.GetElements()
	if err != nil {
		return rt, xerrors.Errorf(": %w", err)
	}

	return rt, nil
}

// NewHeader はWVが出力するCSVのHeader情報をテキストから解析してHeader structに入れて返す
// in: Header Text
// return: Header
func NewHeader(text string) Header {
	var rt Header

	lines := strings.Split(text, "\n")
	rt.Format = strings.Trim(lines[0], "#format ")
	rt.SavedTime = strings.Trim(lines[1], "#[Custom WaveView] saved ")
	rt.Signal = strings.Trim(lines[2], "TIME ,")

	return rt
}

// GetDocument はファイルからテキストを読み出してDocument structにする
// returns: Document, error
func (p WVParser) GetDocument() (Document, error) {
	var rt Document

	b, err := ioutil.ReadFile(p.FilePath)
	if err != nil {
		return Document{}, err
	}
	splited := strings.Split(string(b), "#")

	// Header解析
	rt.Header = NewHeader(splited[1] + splited[2])

	// Data解析
	rt.TextElements = []TextElement{}
	for _, v := range splited[3:] {
		rt.TextElements = append(rt.TextElements, TextElement(v))
	}

	return rt, nil
}

// GetElement はDocumentのTextElementsからElementのスライスを生成して返す
// returns: []*Element, error
func (d Document) GetElements() ([]*Element, error) {
	var rt []*Element

	for _, v := range d.TextElements{
		e,err := v.ElementParse()
		if err != nil {
			return rt, xerrors.Errorf(": %w", err)
		}

		rt=append(rt, &e)
	}

	return rt, nil
}
