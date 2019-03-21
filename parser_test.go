package wvparser

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"testing"
)

func TestWVParser_GentDocument(t *testing.T) {
	home, _ := homedir.Dir()
	p := filepath.Join(home, "TestDir")
	_ = os.MkdirAll(p,0755)

	p = filepath.Join(p, "target.csv")

}
