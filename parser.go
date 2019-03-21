package wvparser

type (
	WVParser struct {
		FilePath string
	}
	WVCsv struct {
		Data []*Element
		Header Header
	}
	Header struct {	}
	Document []TextElement
)

func (wp WVParser) Parse() (WVCsv,error) {
	return WVCsv{},nil
}

// GetDocument はファイルからテキストを読み出してDocument structにする
// returns: Document, error
func (p WVParser) GentDocument() (Document, error) {

	return Document{}, nil
}
