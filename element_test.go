package wvparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextElement_ElementParse(t *testing.T) {
	expect := Element{
		Name: "sweep 0",
		Values: map[float64]float64{
			0.1: 10,
			0.2: 20,
			0.3: 30,
		},
	}

	te := TextElement(`# sweep 0
 0.1 , 10
 0.2 , 20
 0.3 , 30`)

	actual, err := te.ElementParse()

	assert.NoError(t, err)
	assert.Equal(t, expect, actual)
}
