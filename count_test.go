package wvparser

import (
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/unit"
	"path/filepath"
	"testing"
)

func TestParseFilter(t *testing.T) {
	as := assert.New(t)
	t.Run("2.5n:<=0.4", func(t *testing.T) {
		key,st,err := ParseFilter("2.5n:<=:0.4")
		as.NoError(err)
		as.Equal(key,float64(2.5*unit.Nano))
		as.Equal(st,"<=:0.4")
	})
	t.Run("0.7m:>50k", func(t *testing.T) {
		key,st,err := ParseFilter("0.7m:>:50k")
		as.NoError(err)
		as.Equal(key,0.7*unit.Milli)
		as.Equal(st,">:50k")
	})

	t.Run("2.5n:==0.4", func(t *testing.T) {
		key,st,err := ParseFilter("2.5n:==:0.4")
		as.NoError(err)
		as.Equal(key,2.5*unit.Nano)
		as.Equal(st,"==:0.4")
	})
	t.Run("0.7m:!=50k", func(t *testing.T) {
		key,st,err := ParseFilter("0.7m:!=:50k")
		as.NoError(err)
		as.Equal(key,0.7*unit.Milli)
		as.Equal(st,"!=:50k")
	})
}

func TestCounter_AddFilter(t *testing.T) {
	c := Counter{
		Filters: map[float64]string{},
	}

	assert.NoError(t, c.AddFilter("2.5n:>=:0.4"))
	assert.Equal(t,c.Filters[2.5*unit.Nano],">=:0.4")
}

func TestCounter_Aggregate(t *testing.T) {
	c := NewCounter("2.5n:>=:0.4","10n:>=:0.4","17.5n:>=:0.4")
	expect := int64(0)
	csv, err := WVParser{FilePath:filepath.Join("test","target.csv")}.Parse()
	assert.NoError(t,err)
	actual := c.Aggregate(csv)

	assert.Equal(t, expect,actual)
}

func TestCounter_Aggregate2(t *testing.T) {
	c := NewCounter("2.5n:>=:0.4","10n:<=:0.4","17.5n:<=:0.4")
	expect := int64(7)
	csv, err := WVParser{FilePath:filepath.Join("test","target.csv")}.Parse()
	assert.NoError(t,err)
	actual := c.Aggregate(csv)

	assert.Equal(t, expect,actual)

}
